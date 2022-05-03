package akskrequest

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Response struct {
	Metrics []struct {
		Namespace  string `json:"namespace"`
		MetricName string `json:"metric_name"`
		Dimensions []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"dimensions"`
		Datapoints []struct {
			Average   float64 `json:"average"`
			Min       float64 `json:"min"`
			Max       float64 `json:"max"`
			Sum       float64 `json:"sum"`
			Timestamp int64   `json:"timestamp"`
		} `json:"datapoints"`
		Unit string `json:"unit"`
	} `json:"metrics"`
}

type ErrorMsg struct {
	ErrorMessage string `json:"error_msg"`
	ErrorCode    string `json:"error_code"`
	RequestID    string `json:"request_id"`
}

//Calculate value to export in zabbix response
func CalculateValue(responseRequest []byte, filter string) (map[string]float64, error) {
	allMetricValue := make(map[string]float64)
	responseValue := Response{}
	errorMsg := ErrorMsg{}
	result := 0.0
	var total float64

	//Get JSON response in Struct
	json.Unmarshal(responseRequest, &responseValue)

	//If no value => error
	if responseValue.Metrics == nil {
		json.Unmarshal(responseRequest, &errorMsg)
		return nil, fmt.Errorf(strings.Split(errorMsg.ErrorMessage, ", canonical")[0])
	}

	//Loop on each metric
	for _, metric := range responseValue.Metrics {
		total = float64(len(metric.Datapoints))
		//Get value and calculate depending on filter
		for _, point := range metric.Datapoints {
			if filter == "average" {
				if result == -1 {
					result = 0
				}
				result += point.Average
			} else if filter == "min" {
				if result == -1 || result > point.Min {
					result = point.Min
				}
			} else if filter == "max" {
				if result < point.Max {
					result = point.Max
				}
			} else if filter == "sum" {
				if result == -1 {
					result = 0
				}
				result += point.Sum
			}
		}
		if filter == "average" && total != 0 {
			result = result / total
		}
		//Set value for a metric
		allMetricValue[metric.MetricName] = result
	}

	return allMetricValue, nil
}
