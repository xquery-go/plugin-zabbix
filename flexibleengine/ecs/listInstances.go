package ecs

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	ECS []ECSDetail `json:"servers"`
}

type ECSDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]ECSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	ecsDetails := ResultDetails{}

	url := "https://ecs." + region + "." + akskrequest.EndpointDomain + "/v2.1/" + projectID + "/servers/detail"

	response, err := s.MakeRequestGET(projectID, region, "ecs", url)
	if err != nil {
		return ecsDetails.ECS, err
	}

	json.Unmarshal(response, &ecsDetails)

	return ecsDetails.ECS, nil
}
