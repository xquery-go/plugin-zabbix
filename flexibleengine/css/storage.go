package css

// CalculStorage calcul CSS Storage value
func CalculStorage(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
