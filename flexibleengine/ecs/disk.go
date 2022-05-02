package ecs

func CalculDisk(params []string, metric string) (result interface{}, err error) {
	result, err = CheckMetric(params, metric, "AGT.ECS")
	return
}
