package nat

import (
	"errors"
	"fmt"
	"strconv"

	"zabbix.com/pkg/plugin"
	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
	manageresponse "zabbix.com/plugins/flexibleengine/manageResponse"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	if len(params) != 8 {
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
	natID := params[3]
	if natID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
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
	s := akskrequest.Signer{
		Key:    accessKey,
		Secret: secretKey,
	}

	dimension := map[string]interface{}{
		"name":  "nat_gateway_id",
		"value": natID,
	}
	namespace := "SYS.NAT"
	metricsList := []string{"snat_connection"}

	frameInt, _ := strconv.Atoi(frame)

	response, err := s.MakeRequest(projectID, region, frameInt, period, filter, dimension, namespace, metricsList)
	if err != nil {
		return nil, fmt.Errorf("Error in one parameter for make request ($REGION or $PROJECT_ID)")
	}
	value, err := manageresponse.CalculateValue(response, filter)
	if err != nil {
		return nil, err
	}

	if value == -1 {
		return nil, fmt.Errorf("Error in one parameter")
	}

	return value, nil
}

func init() {
	plugin.RegisterMetrics(&impl, "FlexibleEngine",
		"flexibleengine.nat.connections", "Returns connection count.")
}
