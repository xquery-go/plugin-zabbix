package ecs

import (
	"encoding/json"
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

var stateMap = map[int]string{
	0: "pending",
	1: "running",
	2: "paused",
	3: "shutdown",
	4: "crashed",
}

type Response struct {
	Server struct {
		OSEXTSTSPowerState int    `json:"OS-EXT-STS:power_state"`
		Status             string `json:"status"`
	} `json:"server"`
}

func CalculHealth(params []string) (interface{}, error) {
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
	url := "https://ecs." + region + "." + akskrequest.EndpointDomain + "/v2.1/" + projectID + "/servers/" + instanceID

	response, err := s.MakeRequestGET(projectID, region, "ecs", url)
	if err != nil {
		return nil, err
	}
	responseValue := Response{}
	errorMsg := ErrorMsg{}
	json.Unmarshal(response, &responseValue)

	if responseValue.Server.Status == "" {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	value := "state: " + stateMap[responseValue.Server.OSEXTSTSPowerState] + ", Summary: " + responseValue.Server.Status

	return value, nil
}
