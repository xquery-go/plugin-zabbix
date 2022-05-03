package obs

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type GetBucketStorageInfoResult struct {
	XMLName      xml.Name `xml:"GetBucketStorageInfoResult"`
	Text         string   `xml:",chardata"`
	Xmlns        string   `xml:"xmlns,attr"`
	Size         string   `xml:"Size"`
	ObjectNumber string   `xml:"ObjectNumber"`
}

type ErrorMsg struct {
	XMLName    xml.Name `xml:"Error"`
	Text       string   `xml:",chardata"`
	Code       string   `xml:"Code"`
	Message    string   `xml:"Message"`
	RequestId  string   `xml:"RequestId"`
	HostId     string   `xml:"HostId"`
	BucketName string   `xml:"BucketName"`
}

func CalculSize(params []string) (interface{}, error) {
	if len(params) != 5 {
		return nil, errors.New("Wrong parameters.")
	}

	accessKey := params[0]
	if accessKey == "" {
		return nil, fmt.Errorf("Need to specify $ACCESS_KEY option.")
	}
	secretKey := params[1]
	if secretKey == "" {
		return nil, fmt.Errorf("Need to specify $SECRET_KEY option.")
	}
	projectID := params[2]
	if projectID == "" {
		return nil, fmt.Errorf("Need to specify $PROJECT_ID option.")
	}
	region := params[3]
	if region == "" {
		return nil, fmt.Errorf("Need to specify $REGION option.")
	}
	bucketName := params[4]
	if bucketName == "" {
		return nil, fmt.Errorf("Need to specify $BUCKET_NAME option.")
	}

	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}
	url := "https://" + bucketName + ".oss." + akskrequest.EndpointDomain + "?storageinfo"

	response, err := s.MakeRequestGETAWS(projectID, region, "oss", url)
	if err != nil {
		return nil, err
	}

	responseValue := GetBucketStorageInfoResult{}
	errorMsg := ErrorMsg{}
	err = xml.Unmarshal(response, &responseValue)
	if err != nil {
		xml.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.Message)
	}

	value, _ := strconv.Atoi(responseValue.Size)
	return value, nil
}
