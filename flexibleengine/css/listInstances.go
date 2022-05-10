package css

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	CSS []CSSDetail `json:"clusters"`
}

type CSSDetail struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tags []Tag  `json:"tags"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]CSSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	cssDetails := ResultDetails{}

	url := "https://css." + region + "." + akskrequest.EndpointDomain + "/v1.0/" + projectID + "/clusters"

	response, err := s.MakeRequestGET(projectID, region, "css", url)
	if err != nil {
		return cssDetails.CSS, err
	}

	json.Unmarshal(response, &cssDetails)

	return cssDetails.CSS, nil
}
