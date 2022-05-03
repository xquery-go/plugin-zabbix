package eip

// CalculTraffic calcul EIP traffic value
func CalculTraffic(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
