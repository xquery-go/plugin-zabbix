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
	"zabbix.com/plugins/flexibleengine/css"
	"zabbix.com/plugins/flexibleengine/dcs"
	"zabbix.com/plugins/flexibleengine/dds"
	"zabbix.com/plugins/flexibleengine/ecs"
	"zabbix.com/plugins/flexibleengine/eip"
	"zabbix.com/plugins/flexibleengine/elb"
	"zabbix.com/plugins/flexibleengine/evs"
	"zabbix.com/plugins/flexibleengine/nat"
	"zabbix.com/plugins/flexibleengine/obs"
	"zabbix.com/plugins/flexibleengine/rds"
	"zabbix.com/plugins/flexibleengine/sfs"
)

var accessKey string
var secretKey string
var projectID string
var projectName string
var region string
var tokenAPI string
var urlZabbix string
var domainName string

type genericObject struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Engine string   `json:"engine"`
	Role   string   `json:"role"`
	Type   string   `json:"type"`
}

func Discovery(params []string) (interface{}, error) {
	start := time.Now()
	err := verifyParams(params)
	if err != nil {
		return nil, err
	}
	result := "OK"

	//Objects list
	objects := []string{"css", "dcs", "dds", "ecs", "eip", "elb", "evs", "nat", "obs", "rds", "sfs"}
	hostGroupId, _ := hostgroup.GetHostGroupIdWithName(tokenAPI, urlZabbix, domainName)
	if err != nil {
		return nil, err
	}

	for _, object := range objects {
		switch object {
		case "ecs":
			//Get all ECS instances in FE
			listECS, err := ecs.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseECSObjectToGenericObject(listECS)
			val, err := processObject(hostGroupId, listObject, "ECS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "nat":
			//Get all NAT instances in FE
			listNAT, err := nat.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseNATObjectToGenericObject(listNAT)
			val, err := processObject(hostGroupId, listObject, "NAT")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "css":
			//Get all CSS instances in FE
			listCSS, err := css.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseCSSObjectToGenericObject(listCSS)
			val, err := processObject(hostGroupId, listObject, "CSS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "eip":
			//Get all EIP instances in FE
			listEIP, err := eip.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseEIPObjectToGenericObject(listEIP)
			val, err := processObject(hostGroupId, listObject, "EIP")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "elb":
			//Get all ELB instances in FE
			listELB, err := elb.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseELBObjectToGenericObject(listELB)
			val, err := processObject(hostGroupId, listObject, "ELB")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "sfs":
			//Get all SFS instances in FE
			listSFS, err := sfs.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseSFSObjectToGenericObject(listSFS)
			val, err := processObject(hostGroupId, listObject, "SFS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "dcs":
			//Get all DCS instances in FE
			listDCS, err := dcs.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseDCSObjectToGenericObject(listDCS)
			val, err := processObject(hostGroupId, listObject, "DCS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "dds":
			//Get all DDS instances in FE
			listDDS, err := dds.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseDDSObjectToGenericObject(listDDS)
			val, err := processObject(hostGroupId, listObject, "DDS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "evs":
			//Get all EVS instances in FE
			listEVS, err := evs.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseEVSObjectToGenericObject(listEVS)
			val, err := processObject(hostGroupId, listObject, "EVS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "rds":
			//Get all RDS instances in FE
			listRDS, err := rds.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseRDSObjectToGenericObject(listRDS)
			val, err := processObject(hostGroupId, listObject, "RDS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		case "obs":
			//Get all OBS instances in FE
			listOBS, err := obs.ListInstances(accessKey, secretKey, region, projectID)
			if err != nil {
				return "", err
			}
			listObject := parseOBSObjectToGenericObject(listOBS)
			val, err := processObject(hostGroupId, listObject, "OBS")
			if err != nil {
				return nil, err
			}
			result += " " + val
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	return result + " time: " + elapsed.String(), nil
}

//parseCSSObjectToGenericObject permits to parse CSS object to generic object
func parseCSSObjectToGenericObject(listCSS []css.CSSDetail) []genericObject {
	listObject := []genericObject{}
	for _, css := range listCSS {
		tags := []string{}
		for _, tag := range css.Tags {
			tags = append(tags, tag.Key+"="+tag.Value)
		}
		object := genericObject{
			Id:   css.Id,
			Name: css.Name,
			Tags: tags,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseECSObjectToGenericObject permits to parse ECS object to generic object
func parseECSObjectToGenericObject(listECS []ecs.ECSDetail) []genericObject {
	listObject := []genericObject{}
	for _, ecs := range listECS {
		object := genericObject{
			Id:   ecs.Id,
			Name: ecs.Name,
			Tags: ecs.Tags,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseNATObjectToGenericObject permits to parse NAT object to generic object
func parseNATObjectToGenericObject(listNAT []nat.NATDetail) []genericObject {
	listObject := []genericObject{}
	for _, nat := range listNAT {
		object := genericObject{
			Id:   nat.Id,
			Name: nat.Name,
			Tags: nat.Tags,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseEIPObjectToGenericObject permits to parse EIP object to generic object
func parseEIPObjectToGenericObject(listEIP []eip.EIPDetail) []genericObject {
	listObject := []genericObject{}
	for _, eip := range listEIP {
		object := genericObject{
			Id:   eip.Id,
			Name: eip.Name,
			Tags: eip.Tags,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseELBObjectToGenericObject permits to parse ELB object to generic object
func parseELBObjectToGenericObject(listELB []elb.ELBDetail) []genericObject {
	listObject := []genericObject{}
	for _, elb := range listELB {
		object := genericObject{
			Id:   elb.Id,
			Name: elb.Name,
			Tags: elb.Tags,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseSFSObjectToGenericObject permits to parse SFS object to generic object
func parseSFSObjectToGenericObject(listSFS []sfs.SFSDetail) []genericObject {
	listObject := []genericObject{}
	for _, sfs := range listSFS {
		object := genericObject{
			Id:   sfs.Id,
			Name: sfs.Name,
			Tags: sfs.Tags,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseDCSObjectToGenericObject permits to parse DCS object to generic object
func parseDCSObjectToGenericObject(listDCS []dcs.DCSDetail) []genericObject {
	listObject := []genericObject{}
	for _, dcs := range listDCS {
		object := genericObject{
			Id:     dcs.Id,
			Name:   dcs.Name,
			Tags:   dcs.Tags,
			Engine: dcs.Engine,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseDDSObjectToGenericObject permits to parse DDS object to generic object
func parseDDSObjectToGenericObject(listDDS []dds.DDSDetail) []genericObject {
	listObject := []genericObject{}
	for _, dds := range listDDS {
		for _, group := range dds.Groups {
			for _, node := range group.Nodes {
				object := genericObject{
					Id:   node.Id,
					Name: node.Name,
					Tags: dds.Tags,
					Role: node.Role,
				}
				listObject = append(listObject, object)
			}

		}
	}
	return listObject
}

//parseEVSObjectToGenericObject permits to parse EVS object to generic object
func parseEVSObjectToGenericObject(listEVS []evs.EVSDetail) []genericObject {
	listObject := []genericObject{}
	for _, evs := range listEVS {
		for _, attachments := range evs.Attachments {
			tags := []string{}
			id := attachments.ServerId + "-" + strings.Split(attachments.Device, "/dev/")[1]
			for key, value := range evs.Tags {
				tags = append(tags, key+"="+value)
			}

			object := genericObject{
				Id:   id,
				Name: evs.Name,
				Tags: tags,
			}
			listObject = append(listObject, object)

		}
	}
	return listObject
}

//parseRDSObjectToGenericObject permits to parse RDS object to generic object
func parseRDSObjectToGenericObject(listRDS []rds.RDSDetail) []genericObject {
	listObject := []genericObject{}
	for _, rds := range listRDS {
		tags := []string{}
		for _, tag := range rds.Tags {
			tags = append(tags, tag.Key+"="+tag.Value)
		}
		object := genericObject{
			Id:     rds.Id,
			Name:   rds.Name,
			Tags:   tags,
			Engine: strings.ToLower(rds.Datastore.Type),
			Type:   "cluster",
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//parseOBSObjectToGenericObject permits to parse OBS object to generic object
func parseOBSObjectToGenericObject(listOBS []obs.BucketDetail) []genericObject {
	listObject := []genericObject{}
	for _, obs := range listOBS {
		object := genericObject{
			Id:   obs.Name,
			Name: obs.Name,
		}
		listObject = append(listObject, object)
	}
	return listObject
}

//processObject permits to create object list
func processObject(hostGroupId string, listObject []genericObject, typeObject string) (string, error) {
	numberObject := 0

	//Get ID for typeObject template
	templateId, err := template.GetTemplateIdWithName(tokenAPI, urlZabbix, "Cloud-FlexibleEngine-"+typeObject)
	if err != nil {
		return "", err
	}

	//If there is object and no template generate error
	if templateId == "-1" && len(listObject) != 0 {
		return "Template Cloud-FlexibleEngine-" + typeObject + " doesn't exists.", nil
	} else if templateId == "-1" {
		return typeObject + ": " + strconv.Itoa(numberObject), nil
	}

	numberObject = len(listObject)

	//Get existing host with template ID
	listHosts, err := host.GetHostInfo(tokenAPI, urlZabbix, templateId)
	listIndex := []int{}

	//Get all object on instance list FE
	for _, objectFE := range listObject {
		find := false
		for i, hostZabbix := range listHosts {
			for _, macro := range hostZabbix.Macros {
				//If instance already exists in Zabbix
				if (macro.Macro == "{$INSTANCE_ID}" && macro.Value == objectFE.Id) || (macro.Macro == "{$DISK_NAME}" && macro.Value == objectFE.Id) || (macro.Macro == "{$BUCKET_NAME}" && macro.Value == objectFE.Id) {
					find = true
					//Change his name if necessary
					if hostZabbix.Name != typeObject+"_"+objectFE.Name+"_"+objectFE.Id[0:5]+"_"+region {
						name := typeObject + "_" + objectFE.Name + "_" + objectFE.Id[0:5] + "_" + region
						host.UpdateHostName(tokenAPI, urlZabbix, name, hostZabbix.Id)
					}
					//Change his tags
					tags := addTags(objectFE.Tags, typeObject)
					host.UpdateHostTag(tokenAPI, urlZabbix, tags, hostZabbix.Id)
					//Append index in index list to remove
					listIndex = append(listIndex, i)
				}
			}
		}
		//If instance not already exists in Zabbix
		if !find {
			//Set his name, tags and macro and create it
			name := typeObject + "_" + objectFE.Name + "_" + objectFE.Id[0:5] + "_" + region
			tags := addTags(objectFE.Tags, typeObject)
			macros := addMacros(objectFE.Id, typeObject)
			if typeObject == "DCS" {
				macros = append(macros, host.Macro{Macro: "{$ENGINE}", Value: objectFE.Engine, Type: "0"})
			} else if typeObject == "DDS" {
				macros = append(macros, host.Macro{Macro: "{$ROLE}", Value: objectFE.Role, Type: "0"})
			} else if typeObject == "RDS" {
				macros = append(macros, host.Macro{Macro: "{$ENGINE}", Value: objectFE.Engine, Type: "0"})
				macros = append(macros, host.Macro{Macro: "{$TYPE}", Value: objectFE.Type, Type: "0"})
			}
			_ = host.CreateHost(tokenAPI, urlZabbix, name, host.Group{GroupId: hostGroupId}, host.Template{TemplateId: templateId}, tags, macros)
		}
	}

	//Remove instance object already exists in Zabbix and not in FE
	removeExistingObject(listHosts, listIndex)

	return typeObject + ": " + strconv.Itoa(numberObject), nil
}

//removeExistingObject remove objects at particular index
func removeExistingObject(listHosts []host.Host, listIndex []int) {
	//Sort slice to get descending index
	sort.Sort(sort.Reverse(sort.IntSlice(listIndex)))
	//Remove object at the index
	for _, index := range listIndex {
		listHosts = removeIndex(listHosts, index)
	}
	//Delete object if it belong to the project
	for _, hostObject := range listHosts {
		for _, tag := range hostObject.Tags {
			if tag.Tag == "project" && tag.Value == projectName {
				host.DeleteHost(tokenAPI, urlZabbix, hostObject.Id)
			}
		}
	}
}

//addMacros add differents mandatory macros
func addMacros(id string, typeObject string) []host.Macro {
	macros := []host.Macro{}
	macros = append(macros, host.Macro{Macro: "{$ACCESS_KEY}", Value: accessKey, Type: "1"})
	if typeObject == "EVS" {
		macros = append(macros, host.Macro{Macro: "{$DISK_NAME}", Value: id, Type: "0"})
	} else if typeObject == "OBS" {
		macros = append(macros, host.Macro{Macro: "{$BUCKET_NAME}", Value: id, Type: "0"})
	} else {
		macros = append(macros, host.Macro{Macro: "{$INSTANCE_ID}", Value: id, Type: "0"})
	}
	macros = append(macros, host.Macro{Macro: "{$PROJECT_ID}", Value: projectID, Type: "0"})
	macros = append(macros, host.Macro{Macro: "{$REGION}", Value: region, Type: "0"})
	macros = append(macros, host.Macro{Macro: "{$SECRET_KEY}", Value: secretKey, Type: "1"})
	return macros
}

//addTags add FE tags and region, project and type tags to the object
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

//removeIndex remove one element of string at an index
func removeIndex(hosts []host.Host, index int) []host.Host {
	if len(hosts) != 1 {
		return append(hosts[:index], hosts[index+1:]...)
	} else {
		return []host.Host{}
	}
}

// verifyParams verify that all mandatory params are set
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
