package sfs

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	SFS []SFSDetail `json:"shares"`
}

type SFSDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]SFSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	sfsDetails := ResultDetails{}

	url := "https://sfs." + region + "." + akskrequest.EndpointDomain + "/v2/" + projectID + "/shares/detail"

	response, err := s.MakeRequestGET(projectID, region, "sfs", url)
	if err != nil {
		return sfsDetails.SFS, err
	}

	json.Unmarshal(response, &sfsDetails)

	return sfsDetails.SFS, nil
}
