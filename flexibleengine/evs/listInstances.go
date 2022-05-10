package evs

import (
	"encoding/json"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	EVS []EVSDetail `json:"volumes"`
}

type EVSDetail struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Attachments []EVSAttachments  `json:"attachments"`
	Tags        map[string]string `json:"tags"`
}
type EVSAttachments struct {
	ServerId string `json:"server_id"`
	Device   string `json:"device"`
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

	err = json.Unmarshal(response, &evsDetails)
	if err != nil {
		fmt.Println(err.Error())
	}

	return evsDetails.EVS, nil
}
