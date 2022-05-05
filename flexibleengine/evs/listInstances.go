package evs

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	EVS []EVSDetail `json:"volumes"`
}

type EVSDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]EVSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	evsDetails := ResultDetails{}

	url := "https://evs." + region + "." + akskrequest.EndpointDomain + "/v2/" + projectID + "/os-vendor-volumes/detail"

	response, err := s.MakeRequestGET(projectID, region, "evs", url)
	if err != nil {
		return evsDetails.EVS, err
	}

	json.Unmarshal(response, &evsDetails)

	return evsDetails.EVS, nil
}
