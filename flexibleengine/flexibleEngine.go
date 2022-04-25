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
		result, err = ecs.CalculDisk(params, "SlAsH_disk_free")
		return
	case "flexibleengine.ecs.diskused":
		result, err = ecs.CalculDisk(params, "SlAsH_disk_usedPercent")
		return
	case "flexibleengine.ecs.diskread":
		result, err = ecs.CalculDiskIO(params, "disk_read_bytes_rate")
		return
	case "flexibleengine.ecs.diskwrite":
		result, err = ecs.CalculDiskIO(params, "disk_write_bytes_rate")
		return
	case "flexibleengine.ecs.diskrequestread":
		result, err = ecs.CalculDiskIO(params, "disk_read_requests_rate")
		return
	case "flexibleengine.ecs.diskrequestwrite":
		result, err = ecs.CalculDiskIO(params, "disk_write_requests_rate")
		return
	case "flexibleengine.ecs.health":
		result, err = ecs.CalculHealth(params)
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
		"flexibleengine.ecs.diskused", "Returns disk usage.",
		"flexibleengine.ecs.diskread", "Returns disk read bytes rate.",
		"flexibleengine.ecs.diskwrite", "Returns disk write bytes rate.",
		"flexibleengine.ecs.diskrequestread", "Returns disk read ops.",
		"flexibleengine.ecs.diskrequestwrite", "Returns disk write ops.",
		"flexibleengine.ecs.health", "Returns health.")
}
