package sfs

import (
	"encoding/json"
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResponseStatus struct {
	Shares []struct {
		Id     string `json:"id"`
		Status string `json:"status"`
		Name   string `json:"name"`
		Size   string `json:"size"`
	} `json:"shares"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

func CalculSizeUsage(params []string) (interface{}, error) {
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
	region := params[3]
	if region == "" {
		return nil, fmt.Errorf("Need to specify $REGION option.")
	}
	instanceID := params[4]
	if instanceID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}
	url := "https://sfs." + region + "." + akskrequest.EndpointDomain + "/v2/" + projectID + "/shares/detail"

	response, err := s.MakeRequestGET(projectID, region, "sfs", url)
	if err != nil {
		return nil, err
	}
	responseValue := ResponseStatus{}
	errorMsg := ErrorMsg{}
	json.Unmarshal(response, &responseValue)

	if responseValue.Shares == nil {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	var output string

	for _, share := range responseValue.Shares {
		if share.Id == instanceID {
			output = "Name: " + share.Name + " , Size: " + share.Size
			return output, nil
		}
	}

	return nil, fmt.Errorf("No SFS volume found.")
}
