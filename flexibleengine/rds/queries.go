package rds

import (
	"fmt"
)

// CalculQuerie calcul RDS querie value
func CalculQuerie(params []string, metric string) (result interface{}, err error) {
	rdsEngine := params[9]
	if rdsEngine == "" {
		return nil, fmt.Errorf("Need to specify $ENGINE option.")
	} else if rdsEngine != "mysql" {
		return nil, fmt.Errorf("MySQL Engine is only supported.")
	}

	result, err = CheckMetric(params, metric, true)
	return
}
