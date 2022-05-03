package dds

// CalculStorage calcul DDS storage value
func CalculStorage(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
