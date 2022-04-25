package ecs

import (
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculDiskFree(params []string) (result interface{}, err error) {
	ecsID := params[3]
	if ecsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "instance_id",
		"value": ecsID,
	}
	namespace := "AGT.ECS"
	metricsList := []string{"SlAsH_disk_free"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	//valueTruncate := float64(int(value[metricsList[0]]*100)) / 100

	return value[metricsList[0]], nil
}

func CalculDiskUsedPercent(params []string) (result interface{}, err error) {
	ecsID := params[3]
	if ecsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "instance_id",
		"value": ecsID,
	}
	namespace := "AGT.ECS"
	metricsList := []string{"SlAsH_disk_usedPercent"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	//valueTruncate := float64(int(value[metricsList[0]]*100)) / 100

	return value[metricsList[0]], nil
}