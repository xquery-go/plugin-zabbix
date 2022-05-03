package rds

// CalculCPU calcul RDS cpu value
func CalculCPU(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, false)
	return
}
