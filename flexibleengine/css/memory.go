package css

// CalculMemory calcul CSS Memory value
func CalculMemory(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
