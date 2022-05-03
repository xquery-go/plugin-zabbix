package ecs

// CalculMemory calcul ECS memory value
func CalculMemory(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, "AGT.ECS")
	return
}
