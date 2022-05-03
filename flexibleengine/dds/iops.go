package dds

// CalculIOPS calcul DDS IOPS value
func CalculIOPS(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
