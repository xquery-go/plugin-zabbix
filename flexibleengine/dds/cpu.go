package dds

// CalculCPU calcul DDS CPU value
func CalculCPU(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
