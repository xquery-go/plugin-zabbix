package sfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type listSFS struct {
	Values []SFS
}
type SFS struct {
	Name   string
	Status string
}

// CalculStatus calcul SFS status for one or all SFS
func CalculStatus(params []string) (interface{}, error) {
	// Verify params
	if len(params) < 4 {
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
	var instanceID string
	if params[4] != "" {
		instanceID = params[4]
	}

	//Set the AK/SK to sign and authenticate the request.
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}
	//Set url request
	url := "https://sfs." + region + "." + akskrequest.EndpointDomain + "/v2/" + projectID + "/shares/detail"

	response, err := s.MakeRequestGET(projectID, region, "sfs", url)
	if err != nil {
		return nil, err
	}
	responseValue := ResponseStatus{}
	errorMsg := ErrorMsg{}
	//Get JSON response in Struct
	json.Unmarshal(response, &responseValue)

	//If no value => error
	if responseValue.Shares == nil {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	var output string

	creatingVal := 0
	activeVal := 0
	failedVal := 0
	createfailedVal := 0
	deletedVal := 0
	frozenVal := 0
	var listSFS listSFS

	for _, share := range responseValue.Shares {
		if instanceID != "" && share.Id == instanceID {
			switch strings.ToLower(share.Status) {
			case "100":
				creatingVal++
			case "200":
				activeVal++
			case "300":
				failedVal++
			case "303":
				createfailedVal++
			case "400":
				deletedVal++
			case "800":
				frozenVal++
			}
			val := SFS{
				Name:   share.Name,
				Status: share.Status,
			}
			listSFS.Values = append(listSFS.Values, val)
		} else if instanceID == "" {
			switch share.Status {
			case "100":
				creatingVal++
			case "200":
				activeVal++
			case "300":
				failedVal++
			case "303":
				createfailedVal++
			case "400":
				deletedVal++
			case "800":
				frozenVal++
			}
			val := SFS{
				Name:   share.Name,
				Status: share.Status,
			}
			listSFS.Values = append(listSFS.Values, val)
		}
	}

	status := map[string]string{
		"100": "creating",
		"200": "active",
		"300": "failed",
		"303": "create_failed",
		"400": "deleted",
		"800": "frozen",
	}

	output = "Total SFS shares Creating: " + strconv.Itoa(creatingVal) + ", Active: " + strconv.Itoa(activeVal) + ", Failed: " + strconv.Itoa(failedVal) + ", Create_failed: " + strconv.Itoa(createfailedVal) + ", Deleted: " + strconv.Itoa(deletedVal) + ", Frozen: " + strconv.Itoa(frozenVal) + " - "
	for _, SFS := range listSFS.Values {
		output += "Share: " + SFS.Name + " status: " + status[SFS.Status] + " ;"
	}

	return output, nil
}
