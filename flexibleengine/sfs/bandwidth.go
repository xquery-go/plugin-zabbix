package sfs

// CalculBandwith calcul SFS bandwidth value
func CalculBandwith(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric)
	return
}
