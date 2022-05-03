package css

// CalculNodes calcul CSS Nodes value
func CalculNodes(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
