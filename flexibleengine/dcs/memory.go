package dcs

// CalculMemory calcul DCS Memory value
func CalculMemory(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
