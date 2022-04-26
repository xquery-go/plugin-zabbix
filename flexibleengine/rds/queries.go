package rds

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculQuerie(params []string, metric string) (result interface{}, err error) {
	if len(params) != 10 {
		return nil, errors.New("Wrong parameters.")
	}
	rdsID := params[3]
	if rdsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}
	rdsType := params[8]
	if rdsType == "" {
		return nil, fmt.Errorf("Need to specify $TYPE option.")
	}
	rdsEngine := params[9]
	if rdsEngine == "" {
		return nil, fmt.Errorf("Need to specify $ENGINE option.")
	} else if rdsEngine != "mysql" {
		return nil, fmt.Errorf("MySQL Engine is only supported.")
	}

	dimension := map[string]interface{}{
		"name":  "rds_" + rdsType + "_id",
		"value": rdsID,
	}
	namespace := "SYS.RDS"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
