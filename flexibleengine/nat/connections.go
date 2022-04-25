package nat

import (
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

func CalculConnection(params []string) (result interface{}, err error) {
	natID := params[3]
	if natID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	dimension := map[string]interface{}{
		"name":  "nat_gateway_id",
		"value": natID,
	}
	namespace := "SYS.NAT"
	metricsList := []string{"snat_connection"}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return int(value[metricsList[0]]), nil
}
