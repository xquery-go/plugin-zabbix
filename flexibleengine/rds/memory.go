package rds

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculMemory(params []string) (result interface{}, err error) {
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
	}

	var dimensionName string
	if rdsEngine == "mysql" {
		dimensionName = "rds_" + rdsType + "_id"
	} else if rdsEngine == "postgresql" {
		dimensionName = rdsEngine + "_" + rdsType + "_id"
	} else if rdsEngine == "sqlserver" {
		dimensionName = "rds_" + rdsType + "_" + rdsEngine + "id"
	}

	dimension := map[string]interface{}{
		"name":  dimensionName,
		"value": rdsID,
	}
	namespace := "SYS.RDS"
	metricsList := []string{"rds002_mem_util"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
