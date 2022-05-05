package elb

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	ELB []ELBDetail `json:"loadbalancers"`
}

type ELBDetail struct {
	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]ELBDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	elbDetails := ResultDetails{}

	url := "https://vpc." + region + "." + akskrequest.EndpointDomain + "/v2.0/lbaas/loadbalancers"

	response, err := s.MakeRequestGET(projectID, region, "vpc", url)
	if err != nil {
		return elbDetails.ELB, err
	}

	json.Unmarshal(response, &elbDetails)

	return elbDetails.ELB, nil
}
