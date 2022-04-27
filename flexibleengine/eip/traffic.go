package eip

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculTraffic(params []string, metric string) (result interface{}, err error) {
	if len(params) != 8 {
		return nil, errors.New("Wrong parameters.")
	}
	ecsID := params[3]
	if ecsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "publicip_id",
		"value": ecsID,
	}
	namespace := "SYS.VPC"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
