package rds

import (
	"encoding/json"
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type Response struct {
	Instance []struct {
		Status string `json:"status"`
	} `json:"instances"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

//CalculHealth calcul health of RDS
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
	url := "https://rds." + region + "." + akskrequest.EndpointDomain + "/v3/" + projectID + "/instances?id=" + instanceID

	response, err := s.MakeRequestGET(projectID, region, "rds", url)
	if err != nil {
		return nil, err
	}
	responseValue := Response{}
	errorMsg := ErrorMsg{}

	//Get JSON response in Struct
	json.Unmarshal(response, &responseValue)

	//If no value => error
	if len(responseValue.Instance) == 0 {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	value := responseValue.Instance[0].Status

	return value, nil
}
