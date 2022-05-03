package rds

// CalculTransaction calcul RDS transaction value
func CalculTransaction(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, true)
	return
}
