package rds

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

var engineMap = map[string]string{
	"mysql":      "rds007_conn_active_count",
	"postgresql": "rds042_database_connections",
	"sqlserver":  "rds054_db_connections_in_use",
}

func CalculConnection(params []string) (result interface{}, err error) {
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

	dimension := map[string]interface{}{
		"name":  "rds_" + rdsType + "_id",
		"value": rdsID,
	}
	namespace := "SYS.RDS"
	metricsList := []string{engineMap[rdsEngine]}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return int(value[metricsList[0]]), nil
}
