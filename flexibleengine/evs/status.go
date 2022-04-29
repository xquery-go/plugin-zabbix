package evs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResponseStatus struct {
	Volumes []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"volumes"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

type listEVS struct {
	Values []EVS
}
type EVS struct {
	Name   string
	Status string
}

func CalculStatus(params []string) (interface{}, error) {
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
	url := "https://evs." + region + "." + akskrequest.EndpointDomain + "/v2/" + projectID + "/os-vendor-volumes/detail"

	response, err := s.MakeRequestGET(projectID, region, "evs", url)
	if err != nil {
		return nil, err
	}

	responseValue := ResponseStatus{}
	errorMsg := ErrorMsg{}
	json.Unmarshal(response, &responseValue)

	if responseValue.Volumes == nil {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	var output string
	creatingVal := 0
	availableVal := 0
	inuseVal := 0
	errorVal := 0
	attachingVal := 0
	detachingVal := 0
	restoringbackupVal := 0
	backingupVal := 0
	errorrestoringVal := 0
	uploadingVal := 0
	downloadingVal := 0
	extendingVal := 0
	errorextendingVal := 0
	deletingVal := 0
	errordeletingVal := 0
	rollbackingVal := 0
	errorrollbackingVal := 0

	for _, volume := range responseValue.Volumes {
		if instanceID != "" && volume.Id == instanceID {
			switch volume.Status {
			case "creating":
				creatingVal++
			case "available":
				availableVal++
			case "in-use":
				inuseVal++
			case "error":
				errorVal++
			case "attaching":
				attachingVal++
			case "detaching":
				detachingVal++
			case "restoring-backup":
				restoringbackupVal++
			case "backing-up":
				backingupVal++
			case "error_restoring":
				errorrestoringVal++
			case "uploading":
				uploadingVal++
			case "downloading":
				downloadingVal++
			case "extending":
				extendingVal++
			case "error_extending":
				errorextendingVal++
			case "deleting":
				deletingVal++
			case "error_deleting":
				errordeletingVal++
			case "rollbacking":
				rollbackingVal++
			case "error_rollbacking":
				errorrollbackingVal++
			}
		} else if instanceID == "" {
			switch volume.Status {
			case "creating":
				creatingVal++
			case "available":
				availableVal++
			case "in-use":
				inuseVal++
			case "error":
				errorVal++
			case "attaching":
				attachingVal++
			case "detaching":
				detachingVal++
			case "restoring-backup":
				restoringbackupVal++
			case "backing-up":
				backingupVal++
			case "error_restoring":
				errorrestoringVal++
			case "uploading":
				uploadingVal++
			case "downloading":
				downloadingVal++
			case "extending":
				extendingVal++
			case "error_extending":
				errorextendingVal++
			case "deleting":
				deletingVal++
			case "error_deleting":
				errordeletingVal++
			case "rollbacking":
				rollbackingVal++
			case "error_rollbacking":
				errorrollbackingVal++
			}
		}

	}

	output = "Total EVS volumes Creating: " + strconv.Itoa(creatingVal) + ", Available: " + strconv.Itoa(availableVal) + ", In-use: " + strconv.Itoa(inuseVal) +
		", Error: " + strconv.Itoa(errorVal) + ", Attaching: " + strconv.Itoa(attachingVal) + ", Detaching: " + strconv.Itoa(deletingVal) + ", Restoring-backup: " + strconv.Itoa(restoringbackupVal) +
		", Backing-up: " + strconv.Itoa(backingupVal) + ", Error_restoring: " + strconv.Itoa(errorrestoringVal) + ", Uploading: " + strconv.Itoa(uploadingVal) +
		", Downloading: " + strconv.Itoa(downloadingVal) + ", Extending: " + strconv.Itoa(extendingVal) + ", Error_extending: " + strconv.Itoa(errorextendingVal) +
		", Deleting: " + strconv.Itoa(deletingVal) + ", Error_deleting: " + strconv.Itoa(errordeletingVal) + ", Rollbacking: " + strconv.Itoa(rollbackingVal) +
		", Error_rollbacking: " + strconv.Itoa(errorrollbackingVal)

	return output, nil
}
