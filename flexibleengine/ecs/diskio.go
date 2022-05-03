package ecs

// CalculDiskIO calcul ECS diskIO value
func CalculDiskIO(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, "SYS.ECS")
	return
}
