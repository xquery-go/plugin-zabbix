package rds

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculTransaction(params []string) (result interface{}, err error) {
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

	dimension := map[string]interface{}{
		"name":  "rds_" + rdsType + "_id",
		"value": rdsID,
	}
	namespace := "SYS.RDS"
	metricsList := []string{"rds009_tps"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
