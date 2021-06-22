package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	postES()
}

func getEs() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://192.168.20.30:9200/log_operation_admin/_search", nil)
	req.Header.Add("name", "zhaofan")
	req.Header.Add("age", "3")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println(string(body))
	}
}

func postES() {
	urlValues := url.Values{}
	urlValues.Add("name", "{\n  \"query\":{\n  \"bool\": {\n    \"should\": [\n      {\n        \"wildcard\": {\n          \"account\": {\n            \"wildcard\": \"*wygly*\",\n            \"boost\": 1.0\n          }\n        }\n      },\n      {\n        \"wildcard\": {\n          \"client_ip\": {\n            \"wildcard\": \"*GET*\",\n            \"boost\": 1.0\n          }\n        }\n      },\n      {\n        \"match_phrase\": {\n          \"method_desc\": {\n            \"query\": \"wygly\",\n            \"slop\": 0,\n            \"zero_terms_query\": \"NONE\",\n            \"boost\": 1.0\n          }\n        }\n      }\n    ],\n    \"adjust_pure_negative\": true,\n    \"boost\": 1.0\n  }\n}\n}")
	//urlValues.Add("age", "22")
	//client := &http.Client{}
	resp, err := http.PostForm("http://192.168.20.30:9200/log_operation_admin/_search", urlValues)
	//req,_ := http.NewRequest("POST","http://192.168.20.30:9200/log_operation_admin/_search",urlValues)
	//req.Header.Add("name", "zhaofan")
	//req.Header.Add("age", "3")
	//resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
