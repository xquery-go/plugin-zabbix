package rds

import (
	"encoding/json"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResultDetails struct {
	RDS []RDSDetail `json:"instances"`
}

type RDSDetail struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Datastore RDSDatastore `json:"datastore"`
	Tags      []Tag        `json:"tags"`
}

type RDSDatastore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]RDSDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	rdsDetails := ResultDetails{}

	url := "https://rds." + region + "." + akskrequest.EndpointDomain + "/v3/" + projectID + "/instances"

	response, err := s.MakeRequestGET(projectID, region, "rds", url)
	if err != nil {
		return rdsDetails.RDS, err
	}

	json.Unmarshal(response, &rdsDetails)

	return rdsDetails.RDS, nil
}
