// HWS API Gateway Signature
// based on https://github.com/datastream/aws/blob/master/signv4.go
// Copyright (c) 2014, Xianjie

package akskrequest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	BasicDateFormat     = "20060102T150405Z"
	BasicDateTimeFormat = "20060102"
	Algorithm           = "SDK-HMAC-SHA256"
	HeaderXDate         = "X-Sdk-Date"
	HeaderHost          = "host"
	HeaderAuthorization = "Authorization"
	HeaderContentSha256 = "X-Sdk-Content-Sha256"
	EndpointDomain      = "prod-cloud-ocb.orange-business.com"
)

var obcServiceMap = map[string]string{
	"ecs":       "compute",
	"iam":       "identity",
	"cce":       "ccev2.0",
	"evs":       "volumev2",
	"ces":       "cesv1",
	"rds":       "rdsv3",
	"sfs":       "share",
	"css":       "css",
	"workspace": "workspace",
	"dds":       "ddsv3",
	"elb":       "network",
	"vpc":       "vpc",
	"oss":       "s3",
	"dcs":       "dcsv1",
}

func hmacsha256(key []byte, data string) ([]byte, error) {
	h := hmac.New(sha256.New, []byte(key))
	if _, err := h.Write([]byte(data)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Build a CanonicalRequest from a regular request string
//
// CanonicalRequest =
//  HTTPRequestMethod + '\n' +
//  CanonicalURI + '\n' +
//  CanonicalQueryString + '\n' +
//  CanonicalHeaders + '\n' +
//  SignedHeaders + '\n' +
//  HexEncode(Hash(RequestPayload))
func CanonicalRequest(r *http.Request, signedHeaders []string) (string, error) {
	var hexencode string
	var err error
	if hex := r.Header.Get(HeaderContentSha256); hex != "" {
		hexencode = hex
	} else {
		data, err := RequestPayload(r)
		if err != nil {
			return "", err
		}
		hexencode, err = HexEncodeSHA256Hash(data)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", r.Method, CanonicalURI(r), CanonicalQueryString(r), CanonicalHeaders(r, signedHeaders), strings.Join(signedHeaders, ";"), hexencode), err
}

// CanonicalURI returns request uri
func CanonicalURI(r *http.Request) string {
	pattens := strings.Split(r.URL.Path, "/")
	var uri []string
	for _, v := range pattens {
		uri = append(uri, escape(v))
	}
	urlpath := strings.Join(uri, "/")
	if len(urlpath) == 0 || urlpath[len(urlpath)-1] != '/' {
		urlpath = urlpath + "/"
	}
	return urlpath
}

// CanonicalQueryString
func CanonicalQueryString(r *http.Request) string {
	var keys []string
	query := r.URL.Query()
	for key := range query {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var a []string
	for _, key := range keys {
		k := escape(key)
		sort.Strings(query[key])
		for _, v := range query[key] {
			kv := fmt.Sprintf("%s=%s", k, escape(v))
			a = append(a, kv)
		}
	}
	queryStr := strings.Join(a, "&")
	r.URL.RawQuery = queryStr
	return queryStr
}

// CanonicalHeaders
func CanonicalHeaders(r *http.Request, signerHeaders []string) string {
	var a []string
	header := make(map[string][]string)
	for k, v := range r.Header {
		header[strings.ToLower(k)] = v
	}
	for _, key := range signerHeaders {
		value := header[key]
		if strings.EqualFold(key, HeaderHost) {
			value = []string{r.Host}
		}
		sort.Strings(value)
		for _, v := range value {
			a = append(a, key+":"+strings.TrimSpace(v))
		}
	}
	return fmt.Sprintf("%s\n", strings.Join(a, "\n"))
}

// SignedHeaders
func SignedHeaders(r *http.Request) []string {
	var a []string
	for key := range r.Header {
		a = append(a, strings.ToLower(key))
	}
	sort.Strings(a)
	return a
}

// RequestPayload
func RequestPayload(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return []byte(""), nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte(""), err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return b, err
}

// Create a "String to Sign".
func StringToSign(canonicalRequest string, t time.Time, scope string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(canonicalRequest))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n%s\n%s\n%x",
		Algorithm, time.Now().UTC().Format(BasicDateFormat), scope, hash.Sum(nil)), nil
}

// Create the HWS Signature.
func SignStringToSign(stringToSign string, signingKey []byte) (string, error) {
	hm, err := hmacsha256(signingKey, stringToSign)
	return fmt.Sprintf("%x", hm), err
}

func SdkSignKey(date string, region string, serviceType string, secret string) ([]byte, error) {
	ksecret := "SDK" + secret
	kdate, err := hmacsha256([]byte(ksecret), date)
	kregion, err := hmacsha256(kdate, region)
	kservice, err := hmacsha256(kregion, serviceType)
	kapp, err := hmacsha256(kservice, "sdk_request")
	return kapp, err
}

// HexEncodeSHA256Hash returns hexcode of sha256
func HexEncodeSHA256Hash(body []byte) (string, error) {
	hash := sha256.New()
	if body == nil {
		body = []byte("")
	}
	_, err := hash.Write(body)
	return fmt.Sprintf("%x", hash.Sum(nil)), err
}

// Get the finalized value for the "Authorization" header. The signature parameter is the output from SignStringToSign
func AuthHeaderValue(signature, accessKey string, signedHeaders []string, scope string) string {
	return fmt.Sprintf("%s Credential=%s, SignedHeaders=%s, Signature=%s", Algorithm, accessKey+"/"+scope, strings.Join(signedHeaders, ";"), signature)
}

// Signature HWS meta
type Signer struct {
	Key    string
	Secret string
}

// SignRequest set Authorization header
func (s *Signer) Sign(r *http.Request, region string, service string) error {
	var t time.Time
	var err error
	var dt string
	serviceType := obcServiceMap[service]
	scope := time.Now().UTC().Format(BasicDateTimeFormat) + "/" + region + "/" + serviceType + "/sdk_request"
	if dt = r.Header.Get(HeaderXDate); dt != "" {
		t, err = time.Parse(BasicDateFormat, dt)
	}
	if err != nil || dt == "" {
		t = time.Now().UTC()
		r.Header.Set(HeaderXDate, time.Now().UTC().Format(BasicDateFormat))
	}
	signedHeaders := SignedHeaders(r)
	canonicalRequest, err := CanonicalRequest(r, signedHeaders)
	if err != nil {
		return err
	}
	stringToSign, err := StringToSign(canonicalRequest, t, scope)
	if err != nil {
		return err
	}
	signingKey, err := SdkSignKey(time.Now().UTC().Format(BasicDateTimeFormat), region, serviceType, s.Secret)
	if err != nil {
		return err
	}
	// signature, err := SignStringToSign(stringToSign, []byte(signingKey))
	signature, err := SignStringToSign(stringToSign, signingKey)

	if err != nil {
		return err
	}
	authValue := AuthHeaderValue(signature, s.Key, signedHeaders, scope)
	r.Header.Set(HeaderAuthorization, authValue)
	return nil
}

func CreateRequestBody(dimensionName map[string]interface{}, metricsList []string, namespace string, filter string, period string, frame int) []byte {
	var requestBody []byte
	mectrics := make([]map[string]interface{}, 0)
	//Make timestamp
	t := time.Now().UTC()
	endTime := t.Unix() * 1000
	startTime := (t.Add(time.Duration(-frame) * time.Second)).Unix() * 1000

	//Create each metric in list
	for _, metricName := range metricsList {
		dimension := make([]map[string]interface{}, 0)
		dimension = append(dimension, dimensionName)
		mectrics = append(mectrics, map[string]interface{}{
			"namespace":   namespace,
			"metric_name": metricName,
			"dimensions":  dimension,
		})
	}

	//Create final body
	requestBody, _ = json.Marshal(map[string]interface{}{
		"from":    startTime,
		"to":      endTime,
		"period":  period,
		"filter":  filter,
		"metrics": mectrics,
	})
	return requestBody
}

func (s *Signer) MakeRequest(projectID string, region string, frame int, period string, filter string, dimension map[string]interface{}, namespace string, metricsList []string) ([]byte, error) {
	service := "ces"

	//Get body request
	requestBody := CreateRequestBody(dimension, metricsList, namespace, filter, period, frame)

	//Make request with body
	r, _ := http.NewRequest("POST", "https://ces."+region+"."+EndpointDomain+"/V1.0/"+projectID+"/batch-query-metric-data", ioutil.NopCloser(bytes.NewBuffer(requestBody)))

	//Add header parameters
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("X-OpenStack-Nova-API-Version", "2.26")
	r.Header.Add("X-Project-Id", projectID)
	s.Sign(r, region, service)

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


func (s *Signer) MakeRequestGET(projectID string, region string, service string, url string) ([]byte, error) {
	
	//Make request with body
	r, _ := http.NewRequest("GET", url, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))

	//Add header parameters
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("X-OpenStack-Nova-API-Version", "2.26")
	r.Header.Add("X-Project-Id", projectID)
	s.Sign(r, region, service)

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