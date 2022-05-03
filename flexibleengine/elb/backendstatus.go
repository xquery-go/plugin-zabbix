package elb

// CalculBackendStatus calcul ELB backendstatus value
func CalculBackendStatus(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
