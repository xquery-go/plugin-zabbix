package elb

// CalculHTTPRequests calcul ELB httprequest value
func CalculHTTPRequests(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
