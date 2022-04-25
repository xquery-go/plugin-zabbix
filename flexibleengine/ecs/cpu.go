package ecs

import (
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculCPU(params []string) (result interface{}, err error) {
	ecsID := params[3]
	if ecsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "instance_id",
		"value": ecsID,
	}
	namespace := "SYS.ECS"
	metricsList := []string{"cpu_util"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	valueTruncate := float64(int(value[metricsList[0]]*100)) / 100

	return valueTruncate, nil
}