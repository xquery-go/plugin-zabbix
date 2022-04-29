package dcs

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculMemory(params []string, metric string) (result interface{}, err error) {
	if len(params) != 9 {
		return nil, errors.New("Wrong parameters.")
	}
	dcsId := params[3]
	if dcsId == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}
	engine := params[8]
	if engine == "" {
		return nil, fmt.Errorf("Need to specify $ENGINE option.")
	}
	var dimensionName string
	if engine == "redis" {
		dimensionName = "dcs_instance_id"
	} else if engine == "memcached" {
		dimensionName = "dcs_memcached_instance_id"
	}

	dimension := map[string]interface{}{
		"name":  dimensionName,
		"value": dcsId,
	}
	namespace := "SYS.DCS"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
