package dds

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	DDS []DDSDetail `json:"instances"`
}

type DDSDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]DDSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	ddsDetails := ResultDetails{}

	url := "https://dds." + region + "." + akskrequest.EndpointDomain + "/v3/" + projectID + "/instances"

	response, err := s.MakeRequestGET(projectID, region, "dds", url)
	if err != nil {
		return ddsDetails.DDS, err
	}

	json.Unmarshal(response, &ddsDetails)

	return ddsDetails.DDS, nil
}
