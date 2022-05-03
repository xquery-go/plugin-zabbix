package dds

// CalculDiskIO calcul DDS diskIO value
func CalculDiskIO(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
