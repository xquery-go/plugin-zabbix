package rds

// CalculDiskIO calcul RDS diskIO value
func CalculDiskIO(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, false)
	return
}
