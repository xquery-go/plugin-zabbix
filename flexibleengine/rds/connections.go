package rds

import (
	"fmt"
)

var engineMap = map[string]string{
	"mysql":      "rds007_conn_active_count",
	"postgresql": "rds042_database_connections",
	"sqlserver":  "rds054_db_connections_in_use",
}

// CalculConnection calcul RDS connection value
func CalculConnection(params []string) (result interface{}, err error) {
	//Get value engine for post correct metric
	rdsEngine := params[9]
	if rdsEngine == "" {
		return nil, fmt.Errorf("Need to specify $ENGINE option.")
	}
	result, err = CheckMetric(params, engineMap[rdsEngine], true)
	return

}
