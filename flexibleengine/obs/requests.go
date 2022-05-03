package obs

func CalculRequest(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
