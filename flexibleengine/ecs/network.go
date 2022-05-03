package ecs

// CalculNetwork calcul ECS network value
func CalculNetwork(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, "SYS.ECS")
	return
}
