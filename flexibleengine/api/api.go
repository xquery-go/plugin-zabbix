package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MakeRequestPost(url string, method string, tokenAPI string, params interface{}) ([]byte, error) {
	requestBody := CreateRequestBody(method, tokenAPI, params)
	//fmt.Println("\n", string(requestBody))
	r, _ := http.NewRequest("POST", url+"/api_jsonrpc.php", bytes.NewBuffer(requestBody))
	r.Header.Set("Content-type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// CreateRequestBody create body for POST request
func CreateRequestBody(method string, tokenAPI string, params interface{}) []byte {
	var requestBody []byte
	//Create request body
	requestBody, _ = json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"auth":    tokenAPI,
		"id":      1,
		"params":  params,
	})
	return requestBody
}
