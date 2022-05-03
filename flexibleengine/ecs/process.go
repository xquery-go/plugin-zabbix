package ecs

// CalculProcess calcul ECS process value
func CalculProcess(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, "SYS.ECS")
	return
}
