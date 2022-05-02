package evs

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CheckMetric(params []string, metric string) (result interface{}, err error) {
	if len(params) != 8 {
		return nil, errors.New("Wrong parameters.")
	}
	diskName := params[3]
	if diskName == "" {
		return nil, fmt.Errorf("Need to specify $DISK_NAME option.")
	}

	dimension := map[string]interface{}{
		"name":  "disk_name",
		"value": diskName,
	}
	namespace := "SYS.EVS"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
