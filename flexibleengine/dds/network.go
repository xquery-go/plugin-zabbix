package dds

// CalculNetwork calcul DDS network value
func CalculNetwork(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
