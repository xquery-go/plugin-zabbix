package host

import (
	"encoding/json"

	"zabbix.com/plugins/flexibleengine/api"
)

type ParamsHostGet struct {
	Output       []string `json:"output"`
	TemplateIds  string   `json:"templateids"`
	SelectTags   string   `json:"selectTags"`
	SelectMacros string   `json:"selectMacros"`
}

type ResultHost struct {
	Hosts []Host `json:"result"`
}

type Host struct {
	Id     string      `json:"hostid"`
	Name   string      `json:"name"`
	Macros []HostMacro `json:"macros"`
	Tags   []HostTag   `json:"tags"`
}

type HostMacro struct {
	Macro string `json:"macro"`
	Value string `json:"value"`
}

type HostTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func GetHostInfo(tokenAPI string, url string, templateId string) ([]Host, error) {
	params := ParamsHostGet{}
	params.Output = append(params.Output, "hostid")
	params.Output = append(params.Output, "name")
	params.TemplateIds = templateId
	params.SelectMacros = "extend"
	params.SelectTags = "extend"

	body, err := api.MakeRequestPost(url, "host.get", tokenAPI, params)
	if err != nil {
		return nil, err
	}

	result := ResultHost{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Hosts, nil
}
