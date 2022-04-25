package flexibleengine

import (
	"fmt"

	"zabbix.com/pkg/plugin"
	"zabbix.com/plugins/flexibleengine/ecs"
	"zabbix.com/plugins/flexibleengine/nat"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	switch key {
	case "flexibleengine.nat.connections":
		result, err = nat.CalculConnection(params)
		return
	case "flexibleengine.ecs.cpu":
		result, err = ecs.CalculCPU(params)
		return
	case "flexibleengine.ecs.diskfree":
		result, err = ecs.CalculDiskFree(params)
		return
	case "flexibleengine.ecs.diskused":
		result, err = ecs.CalculDiskUsedPercent(params)
		return
	default:
		return nil, fmt.Errorf("Invalid KEY")
	}

}

func init() {
	plugin.RegisterMetrics(&impl, "FlexibleEngineNat",
		"flexibleengine.nat.connections", "Returns connection count.",
		"flexibleengine.ecs.cpu", "Returns CPU value.",
		"flexibleengine.ecs.diskfree", "Returns disk available space.",
		"flexibleengine.ecs.diskused", "Returns disk usage.")
}
