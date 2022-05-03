package rds

// CalculStorage calcul RDS storage value
func CalculStorage(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, false)
	return
}
