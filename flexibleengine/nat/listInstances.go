package nat

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	NAT []NATDetail `json:"nat_gateways"`
}

type NATDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]NATDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	natDetails := ResultDetails{}

	url := "https://nat." + region + "." + akskrequest.EndpointDomain + "/v2.0/nat_gateways"

	response, err := s.MakeRequestGET(projectID, region, "nat", url)
	if err != nil {
		return natDetails.NAT, err
	}

	json.Unmarshal(response, &natDetails)

	return natDetails.NAT, nil
}
