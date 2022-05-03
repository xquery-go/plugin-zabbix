package akskrequest

import (
	"fmt"
	"strconv"
)

//ExecuteProcess execute verification, request and get value
func ExecuteProcess(params []string, dimension map[string]interface{}, namespace string, metricsList []string) (map[string]float64, error) {
	//Verification all mandatory params
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
	region := params[4]
	if region == "" {
		return nil, fmt.Errorf("Need to specify $REGION option.")
	}
	frame := params[5]
	if frame == "" {
		frame = "3600"
	}
	period := params[6]
	if period == "" {
		period = "1"
	}
	filter := params[7]
	if filter == "" {
		filter = "average"
	}

	//Set the AK/SK to sign and authenticate the request.
	s := Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	frameInt, _ := strconv.Atoi(frame)
	//Make the request to API
	response, err := s.MakeRequest(projectID, region, frameInt, period, filter, dimension, namespace, metricsList)
	if err != nil {
		return nil, fmt.Errorf("Error in one parameter for make request ($REGION or $PROJECT_ID)")
	}
	//Calculate result
	value, err := CalculateValue(response, filter)
	if err != nil {
		return nil, err
	}

	return value, nil
}
