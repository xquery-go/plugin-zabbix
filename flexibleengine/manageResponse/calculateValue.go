package manageresponse

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
			Average   int   `json:"average"`
			Min       int   `json:"min"`
			Max       int   `json:"max"`
			Sum       int   `json:"sum"`
			Timestamp int64 `json:"timestamp"`
		} `json:"datapoints"`
		Unit string `json:"unit"`
	} `json:"metrics"`
}

type ErrorMsg struct {
	ErrorMessage string `json:"error_msg"`
	ErrorCode    string `json:"error_code"`
	RequestID    string `json:"request_id"`
}

func CalculateValue(responseRequest []byte, filter string) (int, error) {
	responseValue := Response{}
	errorMsg := ErrorMsg{}
	result := -1
	var total int

	json.Unmarshal(responseRequest, &responseValue)

	if responseValue.Metrics == nil {
		json.Unmarshal(responseRequest, &errorMsg)
		return result, fmt.Errorf(strings.Split(errorMsg.ErrorMessage, ", canonical")[0])
	}

	for _, metric := range responseValue.Metrics {
		total = len(metric.Datapoints)
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
	}

	if filter == "average" && total != 0 {
		result = result / total
	}

	return result, nil
}
