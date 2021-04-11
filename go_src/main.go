package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/4512%20Middle%20Park%20Dr%2C%20San%20Jose%20CA?unitGroup=us&key=WFKDKA26EJQE8RHVEWLP7UV9K")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result_map map[string]interface{}
	//var result_array []map[string]interface{}
	json.Unmarshal([]byte(body), &result_map)
	fmt.Println(result_map["days"])
	var output = result_map["days"].([]map[string]interface{})
	//json.Unmarshal([]byte(output), &result_array)
	fmt.Println(output)
}
