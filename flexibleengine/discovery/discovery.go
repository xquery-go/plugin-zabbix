package discovery

import (
	"fmt"

	"zabbix.com/plugins/flexibleengine/obs"
)

func Discovery(params []string) (interface{}, error) {
	accessKey := ""
	secretKey := ""
	region := ""
	projectID := ""
	//listECS, _ := ecs.ListInstances(accessKey, secretKey, region, projectID)
	//listCSS, _ := css.ListInstances(accessKey, secretKey, region, projectID)
	//listDCS, _ := dcs.ListInstances(accessKey, secretKey, region, projectID)
	//listDDS, _ := dds.ListInstances(accessKey, secretKey, region, projectID)
	//listEIP, _ := eip.ListInstances(accessKey, secretKey, region, projectID)
	//listELB, _ := elb.ListInstances(accessKey, secretKey, region, projectID)
	//listEVS, _ := evs.ListInstances(accessKey, secretKey, region, projectID)
	//listNAT, _ := nat.ListInstances(accessKey, secretKey, region, projectID)
	//listRDS, _ := rds.ListInstances(accessKey, secretKey, region, projectID)
	//listSFS, _ := sfs.ListInstances(accessKey, secretKey, region, projectID)
	listOBS, _ := obs.ListInstances(accessKey, secretKey, region, projectID)
	fmt.Println(listOBS)
	return "OK", nil
}
