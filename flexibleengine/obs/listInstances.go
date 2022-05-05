package obs

import (
	"encoding/xml"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ListAllMyBucketsResult struct {
	XMLName xml.Name `xml:"ListAllMyBucketsResult"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Buckets struct {
		Text   string         `xml:",chardata"`
		Bucket []BucketDetail `xml:"Bucket"`
	} `xml:"Buckets"`
}

type BucketDetail struct {
	Text         string `xml:",chardata"`
	Name         string `xml:"Name"`
	CreationDate string `xml:"CreationDate"`
	BucketType   string `xml:"BucketType"`
}

func ListInstances(accessKey string, secretKey string, region string, projectID string) ([]BucketDetail, error) {
	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	obsDetails := ListAllMyBucketsResult{}

	url := "https://oss." + akskrequest.EndpointDomain

	response, err := s.MakeRequestGETAWS(projectID, region, "oss", url)
	if err != nil {
		return nil, err
	}

	xml.Unmarshal(response, &obsDetails)

	return obsDetails.Buckets.Bucket, nil
}
