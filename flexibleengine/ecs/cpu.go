package ecs

// CalculCPU calcul ECS CPU value
func CalculCPU(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, "SYS.ECS")
	return
}
