package ecs

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculMemory(params []string) (result interface{}, err error) {
	if len(params) != 8 {
		return nil, errors.New("Wrong parameters.")
	}
	ecsID := params[3]
	if ecsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "instance_id",
		"value": ecsID,
	}
	namespace := "AGT.ECS"
	metricsList := []string{"mem_usedPercent"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	valueTruncate := float64(int(value[metricsList[0]]*100)) / 100

	return valueTruncate, nil
}
