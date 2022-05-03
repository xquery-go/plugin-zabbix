package dds

// CalculMemory calcul DDS memory value
func CalculMemory(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
