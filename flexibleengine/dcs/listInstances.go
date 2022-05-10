package dcs

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	DCS []DCSDetail `json:"instances"`
}

type DCSDetail struct {
	Id     string `json:"instance_id"`
	Engine string `json:"engine"`
	Name   string `json:"name"`
	Tags   []string
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]DCSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	dcsDetails := ResultDetails{}

	url := "https://dcs." + region + "." + akskrequest.EndpointDomain + "/v1.0/" + projectID + "/instances"

	response, err := s.MakeRequestGET(projectID, region, "dcs", url)
	if err != nil {
		return dcsDetails.DCS, err
	}

	json.Unmarshal(response, &dcsDetails)

	return dcsDetails.DCS, nil
}
