package obs

// CalculRequest calcul OBS request value
func CalculRequest(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
