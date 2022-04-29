package css

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

type ResponseStatus struct {
	Clusters []struct {
		Id     string `json:"id"`
		Status string `json:"status"`
		Name   string `json:"name"`
	} `json:"clusters"`
}

type ErrorMsg struct {
	ItemNotFound struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"itemNotFound"`
}

type listCSS struct {
	Values []CSS
}
type CSS struct {
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
	url := "https://css." + region + "." + akskrequest.EndpointDomain + "/v1.0/" + projectID + "/clusters"

	response, err := s.MakeRequestGET(projectID, region, "css", url)
	if err != nil {
		return nil, err
	}
	responseValue := ResponseStatus{}
	errorMsg := ErrorMsg{}
	json.Unmarshal(response, &responseValue)

	if responseValue.Clusters == nil {
		json.Unmarshal(response, &errorMsg)
		return nil, fmt.Errorf(errorMsg.ItemNotFound.Message)
	}

	var output string
	availableVal := 0
	createdVal := 0
	unavailableVal := 0
	var listCSS listCSS

	for _, cluster := range responseValue.Clusters {
		if instanceID != "" && cluster.Id == instanceID {
			switch cluster.Status {
			case "200":
				availableVal++
			case "100":
				createdVal++
			case "303":
				unavailableVal++
			}
			val := CSS{
				Name:   cluster.Name,
				Status: cluster.Status,
			}
			listCSS.Values = append(listCSS.Values, val)
		} else if instanceID == "" {
			switch cluster.Status {
			case "200":
				availableVal++
			case "100":
				createdVal++
			case "303":
				unavailableVal++
			}
			val := CSS{
				Name:   cluster.Name,
				Status: cluster.Status,
			}
			listCSS.Values = append(listCSS.Values, val)
		}
	}

	status := map[string]string{
		"200": "available",
		"100": "created",
		"303": "unavailable",
	}

	output = "Total CSS instances Available: " + strconv.Itoa(availableVal) + ", Created: " + strconv.Itoa(createdVal) + ", Unavailable: " + strconv.Itoa(unavailableVal) + " - "
	for _, CSS := range listCSS.Values {
		output += "Server: " + CSS.Name + " status: " + status[CSS.Status] + " ;"
	}

	return output, nil
}
