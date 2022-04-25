package ecs

import (
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculProcess(params []string, metric string) (result interface{}, err error) {
	ecsID := params[3]
	if ecsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "instance_id",
		"value": ecsID,
	}
	namespace := "SYS.ECS"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}
	if value[metricsList[0]] == -1.0 {
		value[metricsList[0]] = 0.0
	}

	return value[metricsList[0]], nil
}
