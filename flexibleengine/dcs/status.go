package dcs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResponseStatus struct {
	Instances []struct {
		Id     string `json:"id"`
		Status string `json:"status"`
		Name   string `json:"name"`
	} `json:"instances"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

type listDCS struct {
	Values []DCS
}
type DCS struct {
	Name   string
	Status string
}

// CalculStatus calcul DCS status for one or all DCS
func CalculStatus(params []string) (interface{}, error) {
	// Verify params
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
	//Set url request
	url := "https://dcs." + region + "." + akskrequest.EndpointDomain + "/v1.0/" + projectID + "/instances"

	response, err := s.MakeRequestGET(projectID, region, "dcs", url)
	if err != nil {
		return nil, err
	}
	responseValue := ResponseStatus{}
	errorMsg := ErrorMsg{}

	//Get JSON response in Struct
	json.Unmarshal(response, &responseValue)

	//If no value => error
	if responseValue.Instances == nil {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	var output string
	creatingVal := 0
	createfailedVal := 0
	runningVal := 0
	errorVal := 0
	restartingVal := 0
	extendingVal := 0
	restoringVal := 0
	var listDCS listDCS

	for _, instance := range responseValue.Instances {
		if instanceID != "" && instance.Id == instanceID {
			switch strings.ToLower(instance.Status) {
			case "creating":
				creatingVal++
			case "createfailed":
				createfailedVal++
			case "running":
				runningVal++
			case "error":
				errorVal++
			case "restarting":
				restartingVal++
			case "extending":
				extendingVal++
			case "restoring":
				restoringVal++
			}
			val := DCS{
				Name:   instance.Name,
				Status: instance.Status,
			}
			listDCS.Values = append(listDCS.Values, val)
		} else if instanceID == "" {
			switch instance.Status {
			case "creating":
				creatingVal++
			case "createfailed":
				createfailedVal++
			case "running":
				runningVal++
			case "error":
				errorVal++
			case "restarting":
				restartingVal++
			case "extending":
				extendingVal++
			case "restoring":
				restoringVal++
			}
			val := DCS{
				Name:   instance.Name,
				Status: instance.Status,
			}
			listDCS.Values = append(listDCS.Values, val)
		}
	}

	output = "Total DCS instances Creating: " + strconv.Itoa(creatingVal) + ", Createfailed: " + strconv.Itoa(createfailedVal) + ", Running: " + strconv.Itoa(runningVal) + ", Error: " + strconv.Itoa(errorVal) + ", Restarting: " + strconv.Itoa(restartingVal) + ", Extending: " + strconv.Itoa(extendingVal) + ", Restoring: " + strconv.Itoa(restoringVal) + " - "
	for _, DCS := range listDCS.Values {
		output += "Instance: " + DCS.Name + " status: " + DCS.Status + " ;"
	}

	return output, nil
}
