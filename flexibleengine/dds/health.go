package dds

import (
	"encoding/json"
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type Response struct {
	Instances []struct {
		Groups []struct {
			Nodes []struct {
				Id     string `json:"id"`
				Role   string `json:"role"`
				Status string `json:"status"`
			} `json:"nodes"`
		} `json:"groups"`
	} `json:"instances"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

//CalculHealth calcul health of DDS
func CalculHealth(params []string) (interface{}, error) {
	//Verify params
	if len(params) != 5 {
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
	region := params[4]
	if region == "" {
		return nil, fmt.Errorf("Need to specify $REGION option.")
	}
	instanceID := params[3]
	if instanceID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}
	//set url request
	url := "https://dds." + region + "." + akskrequest.EndpointDomain + "/v3/" + projectID + "/instances"

	response, err := s.MakeRequestGET(projectID, region, "dds", url)
	if err != nil {
		return nil, err
	}
	responseValue := Response{}
	errorMsg := ErrorMsg{}

	//Get JSON response in Struct
	json.Unmarshal(response, &responseValue)

	//If no value => error
	if len(responseValue.Instances) == 0 {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	for _, instance := range responseValue.Instances {
		for _, groups := range instance.Groups {
			for _, node := range groups.Nodes {
				if node.Id == instanceID {
					return "Status: " + node.Status, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("Instance not found.")
}
