package host

import (
	"zabbix.com/plugins/flexibleengine/api"
)

type ParamsUpdateHostName struct {
	HostId string `json:"hostid"`
	Name   string `json:"name"`
}

type ParamsUpdateHostTags struct {
	HostId string `json:"hostid"`
	Tags   []Tag  `json:"tags"`
}

func UpdateHostName(tokenAPI string, url string, name string, id string) error {
	params := ParamsUpdateHostName{}
	params.HostId = id
	params.Name = name

	_, err := api.MakeRequestPost(url, "host.update", tokenAPI, params)
	if err != nil {
		return err
	}

	return nil
}

func UpdateHostTag(tokenAPI string, url string, tags []Tag, id string) error {
	params := ParamsUpdateHostTags{}
	params.HostId = id
	params.Tags = tags

	_, err := api.MakeRequestPost(url, "host.update", tokenAPI, params)
	if err != nil {
		return err
	}

	return nil
}
