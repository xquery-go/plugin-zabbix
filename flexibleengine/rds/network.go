package rds

// CalculNetwork calcul RDS network value
func CalculNetwork(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, false)
	return
}
