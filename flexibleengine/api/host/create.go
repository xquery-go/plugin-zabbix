package host

import (
	"encoding/json"
	"fmt"

	"zabbix.com/plugins/flexibleengine/api"
)

type ParamsHostCreate struct {
	Host       string      `json:"host"`
	Interfaces []Interface `json:"interfaces"`
	Groups     []Group     `json:"groups"`
	Tags       []Tag       `json:"tags"`
	Templates  []Template  `json:"templates"`
	Macros     []Macro     `json:"macros"`
}

type Interface struct {
	Type  int    `json:"type"`
	Main  int    `json:"main"`
	UseIP int    `json:"useip"`
	Ip    string `json:"ip"`
	Dns   string `json:"dns"`
	Port  string `json:"port"`
}

type Group struct {
	GroupId string `json:"groupid"`
}

type Tag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Template struct {
	TemplateId string `json:"templateid"`
}

type Macro struct {
	Macro string `json:"macro"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type ResultHostCreate struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Result struct {
		HostIds []string `json:"hostids"`
	} `json:"result"`
}

func CreateHost(tokenAPI string, url string, hostName string, group Group, template Template, tags []Tag, macros []Macro) error {
	params := ParamsHostCreate{}
	params.Interfaces = append(params.Interfaces, Interface{Type: 1, Main: 1, UseIP: 1, Ip: "127.0.0.1", Dns: "", Port: "10050"})
	params.Host = hostName
	params.Tags = tags
	params.Macros = macros
	params.Groups = append(params.Groups, group)
	params.Templates = append(params.Templates, template)

	body, err := api.MakeRequestPost(url, "host.create", tokenAPI, params)
	if err != nil {
		return err
	}

	resultHost := ResultHostCreate{}
	json.Unmarshal(body, &resultHost)

	if len(resultHost.Result.HostIds) == 1 {
		return nil
	}
	if resultHost.Error.Message != "" {
		return fmt.Errorf(resultHost.Error.Message)
	}
	return nil

}
