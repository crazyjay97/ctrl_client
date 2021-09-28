package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

var Token string

const baseUrl = "http://127.0.0.1:20219/api/v1/"

func Post(url string, data interface{}) ([]byte, int) {
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(baseUrl+url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return result, resp.StatusCode
}

func Get(url string, params url.Values) ([]byte, int) {
	request, _ := http.NewRequest("GET", baseUrl+url, nil)
	if params != nil {
		query := request.URL.Query()
		for key, _ := range params {
			query.Add(key, params.Get(key))
		}

	}
	request.Header.Set("Authorization", Token)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return result, resp.StatusCode
}
