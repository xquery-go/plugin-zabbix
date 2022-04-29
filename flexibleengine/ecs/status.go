package ecs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResponseStatus struct {
	Servers []struct {
		Id      string `json:"id"`
		VMState string `json:"OS-EXT-STS:vm_state"`
		Name    string `json:"name"`
		Status  string `json:"status"`
	} `json:"servers"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

type listECS struct {
	Values []ECS
}
type ECS struct {
	Name   string
	Status string
}

func CalculStatus(params []string) (interface{}, error) {
	if len(params) < 4 {
		return nil, errors.New("Wrong parameters.")
	}

	accessKey := params[0]
	if accessKey == "" {
		return nil, fmt.Errorf("Need to specify $ACCESS_KEY option.")
	}
	secretKey := params[1]
	if secretKey == "" {
		return nil, fmt.Errorf("Need to specify $SECRET_KEY option.")
	}
	projectID := params[2]
	if projectID == "" {
		return nil, fmt.Errorf("Need to specify $PROJECT_ID option.")
	}
	region := params[3]
	if region == "" {
		return nil, fmt.Errorf("Need to specify $REGION option.")
	}
	var instanceID string
	if params[4] != "" {
		instanceID = params[4]
	}

	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}
	url := "https://ecs." + region + "." + akskrequest.EndpointDomain + "/v2.1/" + projectID + "/servers/detail"

	response, err := s.MakeRequestGET(projectID, region, "ecs", url)
	if err != nil {
		return nil, err
	}
	responseValue := ResponseStatus{}
	errorMsg := ErrorMsg{}
	json.Unmarshal(response, &responseValue)

	if responseValue.Servers == nil {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	var output string
	activeVal := 0
	stoppedVal := 0
	errorVal := 0
	var listECS listECS

	for _, server := range responseValue.Servers {
		if instanceID != "" && server.Id == instanceID {
			switch server.VMState {
			case "active":
				activeVal++
			case "stopped":
				stoppedVal++
			case "error":
				errorVal++
			}
			val := ECS{
				Name:   server.Name,
				Status: server.Status,
			}
			listECS.Values = append(listECS.Values, val)
		} else if instanceID == "" {
			switch server.VMState {
			case "active":
				activeVal++
			case "stopped":
				stoppedVal++
			case "error":
				errorVal++
			}
			val := ECS{
				Name:   server.Name,
				Status: server.Status,
			}
			listECS.Values = append(listECS.Values, val)
		}
	}

	output = "Total ECS servers Active: " + strconv.Itoa(activeVal) + ", Stopped: " + strconv.Itoa(stoppedVal) + ", Error: " + strconv.Itoa(errorVal) + " - "
	for _, ecs := range listECS.Values {
		output += "Server: " + ecs.Name + " status: " + ecs.Status + " ;"
	}

	return output, nil
}
