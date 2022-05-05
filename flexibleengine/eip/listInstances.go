package eip

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	EIP []EIPDetail `json:"publicips"`
}

type EIPDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"bandwidth_name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]EIPDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	eipDetails := ResultDetails{}

	url := "https://vpc." + region + "." + akskrequest.EndpointDomain + "/v1/" + projectID + "/publicips"

	response, err := s.MakeRequestGET(projectID, region, "vpc", url)
	if err != nil {
		return eipDetails.EIP, err
	}

	json.Unmarshal(response, &eipDetails)

	return eipDetails.EIP, nil
}
