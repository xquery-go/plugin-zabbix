package rds

// CalculMemory calcul RDS memory value
func CalculMemory(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, false)
	return
}
