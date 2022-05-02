package dds

func CalculCPU(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
