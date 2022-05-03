package obs

func CalculTraffic(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}