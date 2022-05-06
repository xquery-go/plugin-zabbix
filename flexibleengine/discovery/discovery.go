package discovery

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"zabbix.com/plugins/flexibleengine/api/host"
	hostgroup "zabbix.com/plugins/flexibleengine/api/hostGroup"
	"zabbix.com/plugins/flexibleengine/api/template"
	"zabbix.com/plugins/flexibleengine/ecs"
	"zabbix.com/plugins/flexibleengine/nat"
)

var accessKey string
var secretKey string
var projectID string
var projectName string
var region string
var tokenAPI string
var urlZabbix string
var domainName string

func Discovery(params []string) (interface{}, error) {
	start := time.Now()
	err := verifyParams(params)
	if err != nil {
		return nil, err
	}
	result := "OK"

	objects := []string{"css", "dcs", "dds", "ecs", "eip", "elb", "evs", "nat", "obs", "rds", "sfs"}
	hostGroupId, _ := hostgroup.GetHostGroupIdWithName(tokenAPI, urlZabbix, domainName)
	if err != nil {
		return nil, err
	}

	for _, object := range objects {
		switch object {
		case "ecs":
			val, _ := processObjectECS(hostGroupId)
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "nat":
			val, _ := processObjectNAT(hostGroupId)
			if err != nil {
				return nil, err
			}
			result += " " + val
		}
	}
	//fmt.Println(result)
	//listCSS, _ := css.ListInstances(accessKey, secretKey, region, projectID)
	//listDCS, _ := dcs.ListInstances(accessKey, secretKey, region, projectID)
	//listDDS, _ := dds.ListInstances(accessKey, secretKey, region, projectID)
	//listEIP, _ := eip.ListInstances(accessKey, secretKey, region, projectID)
	//listELB, _ := elb.ListInstances(accessKey, secretKey, region, projectID)
	//listEVS, _ := evs.ListInstances(accessKey, secretKey, region, projectID)
	//listNAT, _ := nat.ListInstances(accessKey, secretKey, region, projectID)
	//listRDS, _ := rds.ListInstances(accessKey, secretKey, region, projectID)
	//listSFS, _ := sfs.ListInstances(accessKey, secretKey, region, projectID)
	//listOBS, _ := obs.ListInstances(accessKey, secretKey, region, projectID)
	t := time.Now()
	elapsed := t.Sub(start)
	return result + "time: " + elapsed.String(), nil
}

func processObjectECS(hostGroupId string) (string, error) {
	numberECS := 0

	templateId, err := template.GetTemplateIdWithName(tokenAPI, urlZabbix, "Cloud-FlexibleEngine-ECS")
	//fmt.Println("\ntemplateID: ", templateId)
	if err != nil {
		return "", err
	}

	listECS, err := ecs.ListInstances(accessKey, secretKey, region, projectID)
	if err != nil {
		return "", err
	}
	//fmt.Println("\nlistECS: ", listECS)

	if templateId == "-1" && len(listECS) != 0 {
		return "Template Cloud-FlexibleEngine-ECS doesn't exists.", nil
	} else if templateId == "-1" {
		return "ECS: " + strconv.Itoa(numberECS), nil
	}

	numberECS = len(listECS)

	listHostsECS, err := host.GetHostInfo(tokenAPI, urlZabbix, templateId)
	//fmt.Println("\nlistHostECS: ", listHostsECS)
	listIndex := []int{}

	for _, ecsFE := range listECS {
		find := false
		for i, hostECS := range listHostsECS {
			for _, macro := range hostECS.Macros {
				if macro.Macro == "{$INSTANCE_ID}" && macro.Value == ecsFE.Id {
					//fmt.Println("\nfind: ", find, ";ecsFE: ", ecsFE, ";hostECS: ", hostECS)
					find = true
					if hostECS.Name != "ecs_"+ecsFE.Name+"_"+ecsFE.Id[0:5]+"_"+region {
						name := "ecs_" + ecsFE.Name + "_" + ecsFE.Id[0:5] + "_" + region
						host.UpdateHostName(tokenAPI, urlZabbix, name, hostECS.Id)
					}

					tags := addTags(ecsFE.Tags, "ecs")
					host.UpdateHostTag(tokenAPI, urlZabbix, tags, hostECS.Id)

					listIndex = append(listIndex, i)
				}
			}
		}
		if !find {
			//fmt.Println("\nfind: ", find, " ;ecsFE: ", ecsFE)
			name := "ecs_" + ecsFE.Name + "_" + ecsFE.Id[0:5] + "_" + region
			tags := addTags(ecsFE.Tags, "ecs")
			macros := addMacros(ecsFE.Id)
			_ = host.CreateHost(tokenAPI, urlZabbix, name, host.Group{GroupId: hostGroupId}, host.Template{TemplateId: templateId}, tags, macros)
		}
	}

	removeExistingObject(listHostsECS, listIndex)

	return "ECS: " + strconv.Itoa(numberECS), nil
}

func processObjectNAT(hostGroupId string) (string, error) {
	numberNAT := 0

	templateId, err := template.GetTemplateIdWithName(tokenAPI, urlZabbix, "Cloud-FlexibleEngine-NAT")
	//fmt.Println("\ntemplateID: ", templateId)
	if err != nil {
		return "", err
	}

	listNAT, err := nat.ListInstances(accessKey, secretKey, region, projectID)
	if err != nil {
		return "", err
	}
	//fmt.Println("\nlistNAT: ", listNAT)

	if templateId == "-1" && len(listNAT) != 0 {
		return "Template Cloud-FlexibleEngine-NAT doesn't exists.", nil
	} else if templateId == "-1" {
		return "NAT: " + strconv.Itoa(numberNAT), nil
	}

	numberNAT = len(listNAT)

	listHostsNAT, err := host.GetHostInfo(tokenAPI, urlZabbix, templateId)
	//fmt.Println("\nlistHostNAT: ", listHostsNAT)
	listIndex := []int{}

	for _, natFE := range listNAT {
		find := false
		for i, hostNAT := range listHostsNAT {
			for _, macro := range hostNAT.Macros {
				if macro.Macro == "{$INSTANCE_ID}" && macro.Value == natFE.Id {
					//fmt.Println("\nfind: ", find, ";natFE: ", natFE, ";hostNAT: ", hostNAT)
					find = true
					if hostNAT.Name != "nat_"+natFE.Name+"_"+natFE.Id[0:5]+"_"+region {
						name := "nat_" + natFE.Name + "_" + natFE.Id[0:5] + "_" + region
						host.UpdateHostName(tokenAPI, urlZabbix, name, hostNAT.Id)
					}

					tags := addTags(natFE.Tags, "nat")
					host.UpdateHostTag(tokenAPI, urlZabbix, tags, hostNAT.Id)

					listIndex = append(listIndex, i)
				}
			}
		}
		if !find {
			//fmt.Println("\nfind: ", find, " ;natFE: ", natFE)
			name := "nat_" + natFE.Name + "_" + natFE.Id[0:5] + "_" + region
			tags := addTags(natFE.Tags, "nat")
			macros := addMacros(natFE.Id)
			_ = host.CreateHost(tokenAPI, urlZabbix, name, host.Group{GroupId: hostGroupId}, host.Template{TemplateId: templateId}, tags, macros)
		}
	}

	removeExistingObject(listHostsNAT, listIndex)

	return "NAT: " + strconv.Itoa(numberNAT), nil
}

func removeExistingObject(listHosts []host.Host, listIndex []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(listIndex)))
	//fmt.Println("\nlistIndex: ", listIndex)
	for _, index := range listIndex {
		listHosts = removeIndex(listHosts, index)
	}
	//fmt.Println("\nlistHostNAT: ", listHosts)
	for _, hostObject := range listHosts {
		for _, tag := range hostObject.Tags {
			if tag.Tag == "project" && tag.Value == projectName {
				host.DeleteHost(tokenAPI, urlZabbix, hostObject.Id)
			}
		}
	}
}

func addMacros(id string) []host.Macro {
	macros := []host.Macro{}
	macros = append(macros, host.Macro{Macro: "{$ACCESS_KEY}", Value: accessKey, Type: "1"})
	macros = append(macros, host.Macro{Macro: "{$INSTANCE_ID}", Value: id, Type: "0"})
	macros = append(macros, host.Macro{Macro: "{$PROJECT_ID}", Value: projectID, Type: "0"})
	macros = append(macros, host.Macro{Macro: "{$REGION}", Value: region, Type: "0"})
	macros = append(macros, host.Macro{Macro: "{$SECRET_KEY}", Value: secretKey, Type: "1"})
	return macros
}

func addTags(tagsECS []string, typeObject string) []host.Tag {
	tags := []host.Tag{}
	for _, tag := range tagsECS {
		tagSplit := strings.Split(tag, "=")
		tags = append(tags, host.Tag{Tag: tagSplit[0], Value: tagSplit[1]})
	}
	tags = append(tags, host.Tag{Tag: "region", Value: region})
	tags = append(tags, host.Tag{Tag: "project", Value: projectName})
	tags = append(tags, host.Tag{Tag: "type", Value: typeObject})
	return tags
}

func removeIndex(hosts []host.Host, index int) []host.Host {
	if len(hosts) != 1 {
		return append(hosts[:index], hosts[index+1:]...)
	} else {
		return []host.Host{}
	}
}

func verifyParams(params []string) error {
	if len(params) != 8 {
		return errors.New("Wrong parameters.")
	}
	accessKey = params[0]
	if accessKey == "" {
		return fmt.Errorf("Need to specify $ACCESS_KEY option.")
	}
	secretKey = params[1]
	if secretKey == "" {
		return fmt.Errorf("Need to specify $SECRET_KEY option.")
	}
	projectID = params[2]
	if projectID == "" {
		return fmt.Errorf("Need to specify $PROJECT_ID option.")
	}
	projectName = params[3]
	if projectName == "" {
		return fmt.Errorf("Need to specify $PROJECT_NAME option.")
	}
	region = params[4]
	if region == "" {
		return fmt.Errorf("Need to specify $REGION option.")
	}
	tokenAPI = params[5]
	if tokenAPI == "" {
		return fmt.Errorf("Need to specify $TOKEN_API option.")
	}
	urlZabbix = params[6]
	if urlZabbix == "" {
		return fmt.Errorf("Need to specify $URL_ZABBIX option.")
	}
	domainName = params[7]
	if domainName == "" {
		return fmt.Errorf("Need to specify $DOMAIN_NAME option.")
	}
	return nil
}
