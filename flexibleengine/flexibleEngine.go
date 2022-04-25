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
	case "flexibleengine.ecs.memory":
		result, err = ecs.CalculMemory(params)
		return
	case "flexibleengine.ecs.networkincominginband":
		result, err = ecs.CalculNetwork(params, "network_incoming_bytes_rate_inband")
		return
	case "flexibleengine.ecs.networkoutgoinginband":
		result, err = ecs.CalculNetwork(params, "network_outgoing_bytes_rate_inband")
		return
	case "flexibleengine.ecs.networkincomingoutband":
		result, err = ecs.CalculNetwork(params, "network_incoming_bytes_aggregate_rate")
		return
	case "flexibleengine.ecs.networkoutgoingoutband":
		result, err = ecs.CalculNetwork(params, "network_outgoing_bytes_aggregate_rate")
		return
	case "flexibleengine.ecs.proctotal":
		result, err = ecs.CalculProcess(params, "proc_total_count")
		return
	case "flexibleengine.ecs.procrunning":
		result, err = ecs.CalculProcess(params, "proc_running_count")
		return
	case "flexibleengine.ecs.proczombie":
		result, err = ecs.CalculProcess(params, "proc_zombie_count")
		return
	case "flexibleengine.ecs.procsleeping":
		result, err = ecs.CalculProcess(params, "proc_sleeping_count")
		return
	case "flexibleengine.ecs.procidle":
		result, err = ecs.CalculProcess(params, "proc_idle_count")
		return
	case "flexibleengine.ecs.status":
		result, err = ecs.CalculStatus(params)
		return
	default:
		return nil, fmt.Errorf("Invalid KEY")
	}

}

func init() {
	plugin.RegisterMetrics(&impl, "FlexibleEngine",
		"flexibleengine.nat.connections", "Returns connection count.",
		"flexibleengine.ecs.cpu", "Returns CPU value.",
		"flexibleengine.ecs.diskfree", "Returns disk available space.",
		"flexibleengine.ecs.diskused", "Returns disk usage.",
		"flexibleengine.ecs.diskread", "Returns disk read bytes rate.",
		"flexibleengine.ecs.diskwrite", "Returns disk write bytes rate.",
		"flexibleengine.ecs.diskrequestread", "Returns disk read ops.",
		"flexibleengine.ecs.diskrequestwrite", "Returns disk write ops.",
		"flexibleengine.ecs.health", "Returns health.",
		"flexibleengine.ecs.memory", "Returns memory used.",
		"flexibleengine.ecs.networkincominginband", "Returns network inband incoming bytes rate.",
		"flexibleengine.ecs.networkoutgoinginband", "Returns network inband outgoing bytes rate.",
		"flexibleengine.ecs.networkincomingoutband", "Returns network outband incoming bytes rate.",
		"flexibleengine.ecs.networkoutgoingoutband", "Returns network outband outgoing bytes rate.",
		"flexibleengine.ecs.proctotal", "Returns total process.",
		"flexibleengine.ecs.procrunning", "Returns running count.",
		"flexibleengine.ecs.proczombie", "Returns zombie process.",
		"flexibleengine.ecs.procsleeping", "Returns sleeping process.",
		"flexibleengine.ecs.procidle", "Returns idle process.",
		"flexibleengine.ecs.status", "Returns status ecs.")
}
