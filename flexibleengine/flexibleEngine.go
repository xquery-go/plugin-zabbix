package flexibleengine

import (
	"fmt"

	"zabbix.com/pkg/plugin"
	"zabbix.com/plugins/flexibleengine/css"
	"zabbix.com/plugins/flexibleengine/dcs"
	"zabbix.com/plugins/flexibleengine/dds"
	"zabbix.com/plugins/flexibleengine/discovery"
	"zabbix.com/plugins/flexibleengine/ecs"
	"zabbix.com/plugins/flexibleengine/eip"
	"zabbix.com/plugins/flexibleengine/elb"
	"zabbix.com/plugins/flexibleengine/evs"
	"zabbix.com/plugins/flexibleengine/nat"
	"zabbix.com/plugins/flexibleengine/obs"
	"zabbix.com/plugins/flexibleengine/rds"
	"zabbix.com/plugins/flexibleengine/sfs"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	switch key {
	case "flexibleengine.nat.connections":
		result, err = nat.CalculConnection(params, "snat_connection")
		return
	case "flexibleengine.ecs.cpu":
		result, err = ecs.CalculCPU(params, "cpu_util")
		return
	case "flexibleengine.ecs.disk.free":
		result, err = ecs.CalculDisk(params, "mountPointPrefix_disk_free")
		return
	case "flexibleengine.ecs.disk.used":
		result, err = ecs.CalculDisk(params, "mountPointPrefix_disk_usedPercent")
		return
	case "flexibleengine.ecs.diskio.read":
		result, err = ecs.CalculDiskIO(params, "disk_read_bytes_rate")
		return
	case "flexibleengine.ecs.diskio.write":
		result, err = ecs.CalculDiskIO(params, "disk_write_bytes_rate")
		return
	case "flexibleengine.ecs.diskio.requestread":
		result, err = ecs.CalculDiskIO(params, "disk_read_requests_rate")
		return
	case "flexibleengine.ecs.diskio.requestwrite":
		result, err = ecs.CalculDiskIO(params, "disk_write_requests_rate")
		return
	case "flexibleengine.ecs.health":
		result, err = ecs.CalculHealth(params)
		return
	case "flexibleengine.ecs.memory":
		result, err = ecs.CalculMemory(params, "mem_usedPercent")
		return
	case "flexibleengine.ecs.network.incominginband":
		result, err = ecs.CalculNetwork(params, "network_incoming_bytes_rate_inband")
		return
	case "flexibleengine.ecs.network.outgoinginband":
		result, err = ecs.CalculNetwork(params, "network_outgoing_bytes_rate_inband")
		return
	case "flexibleengine.ecs.network.incomingoutband":
		result, err = ecs.CalculNetwork(params, "network_incoming_bytes_aggregate_rate")
		return
	case "flexibleengine.ecs.network.outgoingoutband":
		result, err = ecs.CalculNetwork(params, "network_outgoing_bytes_aggregate_rate")
		return
	case "flexibleengine.ecs.proc.total":
		result, err = ecs.CalculProcess(params, "proc_total_count")
		return
	case "flexibleengine.ecs.proc.running":
		result, err = ecs.CalculProcess(params, "proc_running_count")
		return
	case "flexibleengine.ecs.proc.zombie":
		result, err = ecs.CalculProcess(params, "proc_zombie_count")
		return
	case "flexibleengine.ecs.proc.sleeping":
		result, err = ecs.CalculProcess(params, "proc_sleeping_count")
		return
	case "flexibleengine.ecs.proc.idle":
		result, err = ecs.CalculProcess(params, "proc_idle_count")
		return
	case "flexibleengine.ecs.status":
		result, err = ecs.CalculStatus(params)
		return
	case "flexibleengine.rds.connections":
		result, err = rds.CalculConnection(params)
		return
	case "flexibleengine.rds.cpu":
		result, err = rds.CalculCPU(params, "rds001_cpu_util")
		return
	case "flexibleengine.rds.diskio.read":
		result, err = rds.CalculDiskIO(params, "rds049_disk_read_throughput")
		return
	case "flexibleengine.rds.diskio.write":
		result, err = rds.CalculDiskIO(params, "rds050_disk_write_throughput")
		return
	case "flexibleengine.rds.health":
		result, err = rds.CalculHealth(params)
		return
	case "flexibleengine.rds.memory":
		result, err = rds.CalculMemory(params, "rds002_mem_util")
		return
	case "flexibleengine.rds.network.input":
		result, err = rds.CalculNetwork(params, "rds004_bytes_in")
		return
	case "flexibleengine.rds.network.output":
		result, err = rds.CalculNetwork(params, "rds005_bytes_out")
		return
	case "flexibleengine.rds.storage":
		result, err = rds.CalculStorage(params, "rds039_disk_util")
		return
	case "flexibleengine.rds.transaction":
		result, err = rds.CalculTransaction(params, "rds009_tps")
		return
	case "flexibleengine.rds.querie":
		result, err = rds.CalculQuerie(params, "rds008_qps")
		return
	case "flexibleengine.rds.querie.delete":
		result, err = rds.CalculQuerie(params, "rds028_comdml_del_count")
		return
	case "flexibleengine.rds.querie.insert":
		result, err = rds.CalculQuerie(params, "rds029_comdml_ins_count")
		return
	case "flexibleengine.rds.querie.insertselect":
		result, err = rds.CalculQuerie(params, "rds030_comdml_ins_sel_count")
		return
	case "flexibleengine.rds.querie.replace":
		result, err = rds.CalculQuerie(params, "rds031_comdml_rep_count")
		return
	case "flexibleengine.rds.querie.replaceselection":
		result, err = rds.CalculQuerie(params, "rds032_comdml_rep_sel_count")
		return
	case "flexibleengine.rds.querie.select":
		result, err = rds.CalculQuerie(params, "rds033_comdml_sel_count")
		return
	case "flexibleengine.rds.querie.update":
		result, err = rds.CalculQuerie(params, "rds034_comdml_upd_count")
		return
	case "flexibleengine.evs.diskio.read":
		result, err = evs.CalculDiskIO(params, "disk_device_read_bytes_rate")
		return
	case "flexibleengine.evs.diskio.write":
		result, err = evs.CalculDiskIO(params, "disk_device_write_bytes_rate")
		return
	case "flexibleengine.evs.diskio.requestread":
		result, err = evs.CalculDiskIO(params, "disk_device_read_requests_rate")
		return
	case "flexibleengine.evs.diskio.requestwrite":
		result, err = evs.CalculDiskIO(params, "disk_device_write_requests_rate")
		return
	case "flexibleengine.evs.diskio.readoperation":
		result, err = evs.CalculDiskIO(params, "disk_device_read_bytes_per_operation")
		return
	case "flexibleengine.evs.diskio.readawait":
		result, err = evs.CalculDiskIO(params, "disk_device_read_await")
		return
	case "flexibleengine.evs.diskio.queuelength":
		result, err = evs.CalculDiskIO(params, "disk_device_queue_length")
		return
	case "flexibleengine.evs.diskio.ioutil":
		result, err = evs.CalculDiskIO(params, "disk_device_io_util")
		return
	case "flexibleengine.evs.diskio.iosvctm":
		result, err = evs.CalculDiskIO(params, "disk_device_io_svctm")
		return
	case "flexibleengine.evs.status":
		result, err = evs.CalculStatus(params)
		return
	case "flexibleengine.eip.traffic.downstream":
		result, err = eip.CalculTraffic(params, "downstream_bandwidth")
		return
	case "flexibleengine.eip.traffic.upstream":
		result, err = eip.CalculTraffic(params, "upstream_bandwidth")
		return
	case "flexibleengine.elb.backendstatus.anormal":
		result, err = elb.CalculBackendStatus(params, "m9_abnormal_servers")
		return
	case "flexibleengine.elb.backendstatus.normal":
		result, err = elb.CalculBackendStatus(params, "ma_normal_servers")
		return
	case "flexibleengine.elb.connection.concurrent":
		result, err = elb.CalculConnection(params, "m1_cps")
		return
	case "flexibleengine.elb.connection.active":
		result, err = elb.CalculConnection(params, "m2_act_conn")
		return
	case "flexibleengine.elb.connection.inactive":
		result, err = elb.CalculConnection(params, "m3_inact_conn")
		return
	case "flexibleengine.elb.httprequests.layer":
		result, err = elb.CalculHTTPRequests(params, "mb_l7_qps")
		return
	case "flexibleengine.elb.httprequests.2xxcodes":
		result, err = elb.CalculHTTPRequests(params, "mc_l7_http_2xx")
		return
	case "flexibleengine.elb.httprequests.3xxcodes":
		result, err = elb.CalculHTTPRequests(params, "md_l7_http_3xx")
		return
	case "flexibleengine.elb.httprequests.4xxcodes":
		result, err = elb.CalculHTTPRequests(params, "me_l7_http_4xx")
		return
	case "flexibleengine.elb.httprequests.5xxcodes":
		result, err = elb.CalculHTTPRequests(params, "mf_l7_http_5xx")
		return
	case "flexibleengine.elb.httprequests.otherstatus":
		result, err = elb.CalculHTTPRequests(params, "m10_l7_http_other_status")
		return
	case "flexibleengine.elb.httprequests.404":
		result, err = elb.CalculHTTPRequests(params, "m11_l7_http_404")
		return
	case "flexibleengine.elb.httprequests.499":
		result, err = elb.CalculHTTPRequests(params, "m12_l7_http_499")
		return
	case "flexibleengine.elb.httprequests.502":
		result, err = elb.CalculHTTPRequests(params, "m13_l7_http_502")
		return
	case "flexibleengine.elb.httprequests.averagelayer":
		result, err = elb.CalculHTTPRequests(params, "m14_l7_rt")
		return
	case "flexibleengine.elb.traffic.incoming":
		result, err = elb.CalculTraffic(params, "m5_in_pps")
		return
	case "flexibleengine.elb.traffic.outgoing":
		result, err = elb.CalculTraffic(params, "m6_out_pps")
		return
	case "flexibleengine.elb.traffic.inbound":
		result, err = elb.CalculTraffic(params, "m7_in_Bps")
		return
	case "flexibleengine.elb.traffic.outbound":
		result, err = elb.CalculTraffic(params, "m8_out_Bps")
		return
	case "flexibleengine.elb.health":
		result, err = elb.CalculHealth(params)
		return
	case "flexibleengine.dds.cpu":
		result, err = dds.CalculCPU(params, "mongo031_cpu_usage")
		return
	case "flexibleengine.dds.diskio.read":
		result, err = dds.CalculDiskIO(params, "mongo037_read_throughput")
		return
	case "flexibleengine.dds.diskio.write":
		result, err = dds.CalculDiskIO(params, "mongo038_write_throughput")
		return
	case "flexibleengine.dds.iops":
		result, err = dds.CalculIOPS(params, "mongo036_iops")
		return
	case "flexibleengine.dds.memory":
		result, err = dds.CalculMemory(params, "mongo032_mem_usage")
		return
	case "flexibleengine.dds.network.out":
		result, err = dds.CalculNetwork(params, "mongo033_bytes_out")
		return
	case "flexibleengine.dds.network.in":
		result, err = dds.CalculNetwork(params, "mongo034_bytes_in")
		return
	case "flexibleengine.dds.storage":
		result, err = dds.CalculStorage(params, "mongo035_disk_usage")
		return
	case "flexibleengine.dds.health":
		result, err = dds.CalculHealth(params)
		return
	case "flexibleengine.css.cpu":
		result, err = css.CalculCPU(params, "max_cpu_usage")
		return
	case "flexibleengine.css.indices.doccount":
		result, err = css.CalculIndices(params, "docs_count")
		return
	case "flexibleengine.css.indices.delete":
		result, err = css.CalculIndices(params, "docs_deleted_count")
		return
	case "flexibleengine.css.indices.count":
		result, err = css.CalculIndices(params, "indices_count")
		return
	case "flexibleengine.css.indices.totalshards":
		result, err = css.CalculIndices(params, "total_shards_count")
		return
	case "flexibleengine.css.indices.primaryshards":
		result, err = css.CalculIndices(params, "primary_shards_count")
		return
	case "flexibleengine.css.memory":
		result, err = css.CalculMemory(params, "max_jvm_heap_usage")
		return
	case "flexibleengine.css.nodes.count":
		result, err = css.CalculNodes(params, "nodes_count")
		return
	case "flexibleengine.css.nodes.data":
		result, err = css.CalculNodes(params, "data_nodes_count")
		return
	case "flexibleengine.css.nodes.coordinating":
		result, err = css.CalculNodes(params, "coordinating_nodes_count")
		return
	case "flexibleengine.css.nodes.master":
		result, err = css.CalculNodes(params, "master_nodes_count")
		return
	case "flexibleengine.css.nodes.ingest":
		result, err = css.CalculNodes(params, "ingest_nodes_count")
		return
	case "flexibleengine.css.storage.total":
		result, err = css.CalculStorage(params, "total_fs_size")
		return
	case "flexibleengine.css.storage.free":
		result, err = css.CalculStorage(params, "free_fs_size")
		return
	case "flexibleengine.css.status":
		result, err = css.CalculStatus(params)
		return
	case "flexibleengine.dcs.memory":
		result, err = dcs.CalculMemory(params, "memory_usage")
		return
	case "flexibleengine.dcs.status":
		result, err = dcs.CalculStatus(params)
		return
	case "flexibleengine.sfs.bandwidth.read":
		result, err = sfs.CalculBandwith(params, "read_bandwidth")
		return
	case "flexibleengine.sfs.bandwidth.write":
		result, err = sfs.CalculBandwith(params, "write_bandwidth")
		return
	case "flexibleengine.sfs.bandwidth.disk":
		result, err = sfs.CalculBandwith(params, "rw_bandwidth")
		return
	case "flexibleengine.sfs.sizeusage":
		result, err = sfs.CalculSizeUsage(params)
		return
	case "flexibleengine.sfs.status":
		result, err = sfs.CalculStatus(params)
		return
	case "flexibleengine.obs.traffic.download":
		result, err = obs.CalculTraffic(params, "download_bytes")
		return
	case "flexibleengine.obs.traffic.upload":
		result, err = obs.CalculTraffic(params, "upload_bytes")
		return
	case "flexibleengine.obs.requests.get":
		result, err = obs.CalculRequest(params, "get_request_count")
		return
	case "flexibleengine.obs.requests.put":
		result, err = obs.CalculRequest(params, "put_request_count")
		return
	case "flexibleengine.obs.requests.4xxcodes":
		result, err = obs.CalculRequest(params, "request_count_4xx")
		return
	case "flexibleengine.obs.requests.5xxcodes":
		result, err = obs.CalculRequest(params, "request_count_5xx")
		return
	case "flexibleengine.obs.size":
		result, err = obs.CalculSize(params)
		return
	case "flexibleengine.discovery":
		result, err = discovery.Discovery(params)
		return
	default:
		return nil, fmt.Errorf("Invalid KEY")
	}
}

func init() {
	plugin.RegisterMetrics(&impl, "FlexibleEngine",
		"flexibleengine.nat.connections", "Returns connection count.",
		"flexibleengine.ecs.cpu", "Returns CPU value.",
		"flexibleengine.ecs.disk.free", "Returns disk available space.",
		"flexibleengine.ecs.disk.used", "Returns disk usage.",
		"flexibleengine.ecs.diskio.read", "Returns disk read bytes rate.",
		"flexibleengine.ecs.diskio.write", "Returns disk write bytes rate.",
		"flexibleengine.ecs.diskio.requestread", "Returns disk read ops.",
		"flexibleengine.ecs.diskio.requestwrite", "Returns disk write ops.",
		"flexibleengine.ecs.health", "Returns health.",
		"flexibleengine.ecs.memory", "Returns memory used.",
		"flexibleengine.ecs.network.incominginband", "Returns network inband incoming bytes rate.",
		"flexibleengine.ecs.network.outgoinginband", "Returns network inband outgoing bytes rate.",
		"flexibleengine.ecs.network.incomingoutband", "Returns network outband incoming bytes rate.",
		"flexibleengine.ecs.network.outgoingoutband", "Returns network outband outgoing bytes rate.",
		"flexibleengine.ecs.proc.total", "Returns total process.",
		"flexibleengine.ecs.proc.running", "Returns running count.",
		"flexibleengine.ecs.proc.zombie", "Returns zombie process.",
		"flexibleengine.ecs.proc.sleeping", "Returns sleeping process.",
		"flexibleengine.ecs.proc.idle", "Returns idle process.",
		"flexibleengine.ecs.status", "Returns status ecs.",
		"flexibleengine.rds.connections", "Returns connection count.",
		"flexibleengine.rds.cpu", "Returns CPU value.",
		"flexibleengine.rds.diskio.read", "Returns disk read throughput.",
		"flexibleengine.rds.diskio.write", "Returns disk write throughput.",
		"flexibleengine.rds.health", "Returns health.",
		"flexibleengine.rds.memory", "Returns memory used.",
		"flexibleengine.rds.network.input", "Returns network input throughput.",
		"flexibleengine.rds.network.output", "Returns network output throughput.",
		"flexibleengine.rds.storage", "Returns storage utilization.",
		"flexibleengine.rds.transaction", "Returns transactions per second.",
		"flexibleengine.rds.querie", "Returns queries per seconde.",
		"flexibleengine.rds.querie.delete", "Returns delete statements per second.",
		"flexibleengine.rds.querie.insert", "Returns insert statements per second.",
		"flexibleengine.rds.querie.insertselect", "Returns insert/select statements per second.",
		"flexibleengine.rds.querie.replace", "Returns replace statements per second.",
		"flexibleengine.rds.querie.replaceselection", "Returns replace_selection statements per second.",
		"flexibleengine.rds.querie.select", "Returns select statements per second.",
		"flexibleengine.rds.querie.update", "Returns update statements per second.",
		"flexibleengine.evs.diskio.read", "Returns disk read bytes rate.",
		"flexibleengine.evs.diskio.write", "Returns disk write bytes rate.",
		"flexibleengine.evs.diskio.requestread", "Returns disk read ops.",
		"flexibleengine.evs.diskio.requestwrite", "Returns disk write ops.",
		"flexibleengine.evs.diskio.readoperation", "Returns avg disk read bytes per operation.",
		"flexibleengine.evs.diskio.readawait", "Returns disk read await.",
		"flexibleengine.evs.diskio.queuelength", "Returns average queue length.",
		"flexibleengine.evs.diskio.ioutil", "Returns disk I/O utilization.",
		"flexibleengine.evs.diskio.iosvctm", "Returns disk I/O service time.",
		"flexibleengine.evs.status", "Returns status evs.",
		"flexibleengine.eip.traffic.downstream", "Returns downstream bandwith.",
		"flexibleengine.eip.traffic.upstream", "Returns upstream bandwith.",
		"flexibleengine.elb.backendstatus.anormal", "Returns unhealthy servers.",
		"flexibleengine.elb.backendstatus.normal", "Returns healthy servers.",
		"flexibleengine.elb.connection.concurrent", "Returns concurrent connections.",
		"flexibleengine.elb.connection.active", "Returns active connections.",
		"flexibleengine.elb.connection.inactive", "Returns inactive connections.",
		"flexibleengine.elb.httprequests.layer", "Returns layer-7 query rate.",
		"flexibleengine.elb.httprequests.2xxcodes", "Returns 2xx status code.",
		"flexibleengine.elb.httprequests.3xxcodes", "Returns 3xx status code.",
		"flexibleengine.elb.httprequests.4xxcodes", "Returns 4xx status code.",
		"flexibleengine.elb.httprequests.5xxcodes", "Returns 5xx status code.",
		"flexibleengine.elb.httprequests.otherstatus", "Returns other status code.",
		"flexibleengine.elb.httprequests.404", "Returns 404 status code.",
		"flexibleengine.elb.httprequests.499", "Returns 499 status code.",
		"flexibleengine.elb.httprequests.502", "Returns 502 status code.",
		"flexibleengine.elb.httprequests.averagelayer", "Returns average layer-7 response time.",
		"flexibleengine.elb.traffic.incoming", "Returns incoming packets rate.",
		"flexibleengine.elb.traffic.outgoing", "Returns outgoing packets rate.",
		"flexibleengine.elb.traffic.inbound", "Returns inbound rate.",
		"flexibleengine.elb.traffic.outbound", "Returns outbound rate.",
		"flexibleengine.elb.health", "Returns health.",
		"flexibleengine.dds.cpu", "Returns CPU utilization.",
		"flexibleengine.dds.diskio.read", "Returns disk read throughput.",
		"flexibleengine.dds.diskio.write", "Returns disk write throughput.",
		"flexibleengine.dds.iops", "Returns I/O per second.",
		"flexibleengine.dds.memory", "Returns memory utilization.",
		"flexibleengine.dds.network.out", "Returns network output throughput.",
		"flexibleengine.dds.network.in", "Returns network input throughput.",
		"flexibleengine.dds.storage", "Returns storage utilization.",
		"flexibleengine.dds.health", "Returns health.",
		"flexibleengine.css.cpu", "Returns max CPU usage.",
		"flexibleengine.css.indices.doccount", "Returns documents count.",
		"flexibleengine.css.indices.delete", "Returns deleted documents count.",
		"flexibleengine.css.indices.count", "Returns indices count.",
		"flexibleengine.css.indices.totalshards", "Returns total shards count.",
		"flexibleengine.css.indices.primaryshards", "Returns primary shards count.",
		"flexibleengine.css.memory", "Returns max JVM heap usage.",
		"flexibleengine.css.nodes.count", "Returns nodes count.",
		"flexibleengine.css.nodes.data", "Returns data nodes count.",
		"flexibleengine.css.nodes.coordinating", "Returns coordination nodes.",
		"flexibleengine.css.nodes.master", "Returns master nodes.",
		"flexibleengine.css.nodes.ingest", "Returns client nodes.",
		"flexibleengine.css.storage.total", "Returns total storage.",
		"flexibleengine.css.storage.free", "Returns free storage.",
		"flexibleengine.css.status", "Returns status css.",
		"flexibleengine.dcs.memory", "Returns memory usage.",
		"flexibleengine.dcs.status", "Returns status dcs.",
		"flexibleengine.sfs.bandwidth.read", "Returns read bandwidth.",
		"flexibleengine.sfs.bandwidth.write", "Returns write bandwidth.",
		"flexibleengine.sfs.bandwidth.disk", "Returns disk write ops.",
		"flexibleengine.sfs.sizeusage", "Returns size usage.",
		"flexibleengine.sfs.status", "Returns status sfs.",
		"flexibleengine.obs.traffic.download", "Returns bytes download.",
		"flexibleengine.obs.traffic.upload", "Returns bytes upload.",
		"flexibleengine.obs.requests.get", "Returns get requests.",
		"flexibleengine.obs.requests.put", "Returns put requests.",
		"flexibleengine.obs.requests.4xxcodes", "Returns 4xx errors.",
		"flexibleengine.obs.requests.5xxcodes", "Returns 5xx errors.",
		"flexibleengine.obs.size", "Returns size of bucket.",
		"flexibleengine.discovery", "Returns all FE instances discovery.")
}
