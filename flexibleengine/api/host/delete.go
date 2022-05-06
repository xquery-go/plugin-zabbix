package host

import "zabbix.com/plugins/flexibleengine/api"

func DeleteHost(tokenAPI string, url string, hostID string) error {
	params := []string{hostID}

	_, err := api.MakeRequestPost(url, "host.delete", tokenAPI, params)
	if err != nil {
		return err
	}

	return nil
}
