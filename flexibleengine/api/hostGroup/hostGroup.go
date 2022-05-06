package hostgroup

import (
	"encoding/json"
	"fmt"
	"strings"

	"zabbix.com/plugins/flexibleengine/api"
)

type ParamsHostGroup struct {
	Filter struct {
		Name []string `json:"name"`
	} `json:"filter"`
}

type ResultHostGroup struct {
	Result []struct {
		Id   string `json:"groupid"`
		Name string `json:"name"`
	} `json:"result"`
}

func GetHostGroupIdWithName(tokenAPI string, url string, hostGroupName string) (string, error) {
	params := ParamsHostGroup{}
	params.Filter.Name = append(params.Filter.Name, hostGroupName)
	params.Filter.Name = append(params.Filter.Name, "Zabbix servers")

	body, err := api.MakeRequestPost(url, "hostgroup.get", tokenAPI, params)
	if err != nil {
		return "", err
	}

	result := ResultHostGroup{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	for _, template := range result.Result {
		if strings.ToLower(hostGroupName) == strings.ToLower(template.Name) {
			return template.Id, nil
		}
	}

	for _, template := range result.Result {
		if "zabbix servers" == strings.ToLower(template.Name) {
			return template.Id, nil
		}
	}

	return "", fmt.Errorf("No host group name")
}
