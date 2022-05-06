package template

import (
	"encoding/json"

	"zabbix.com/plugins/flexibleengine/api"
)

type ParamsTemplate struct {
	Output []string `json:"output"`
	Filter struct {
		Name []string `json:"name"`
	} `json:"filter"`
}

type ResultTemplate struct {
	Result []struct {
		Id string `json:"templateid"`
	} `json:"result"`
}

func GetTemplateIdWithName(tokenAPI string, url string, templateName string) (string, error) {
	params := ParamsTemplate{}
	params.Filter.Name = append(params.Filter.Name, templateName)
	params.Output = append(params.Output, "templateid")

	body, err := api.MakeRequestPost(url, "template.get", tokenAPI, params)
	if err != nil {
		return "", err
	}

	result := ResultTemplate{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result.Result) == 0 {
		return "-1", nil
	}

	return result.Result[0].Id, nil
}
