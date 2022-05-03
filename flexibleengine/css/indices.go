package css

// CalculIndices calcul CSS indices value
func CalculIndices(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
