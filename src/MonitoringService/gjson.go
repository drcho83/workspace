package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//config 파일 읽기
	config_data, err := ioutil.ReadFile("config.json") //--> byte로 받아 오기 때문에 String으로 형 변황을 해줘야 함
	errCheck(err)
	env_data, err := ioutil.ReadFile("env.json") //--> byte로 받아 오기 때문에 String으로 형 변황을 해줘야 함
	errCheck(err)

	var config_data_result map[string]interface{}
	json.Unmarshal([]byte(config_data), &config_data_result)

	var env_data_result map[string]interface{}
	json.Unmarshal([]byte(env_data), &env_data_result)

	for k, v := range config_data_result {
		fmt.Println(k, v)
	}

	fmt.Println("======================")
	fmt.Println(env_data_result["InfluxDB_Address"])
	fmt.Println(env_data_result["Interval"])
	fmt.Println(env_data_result["Stored_DB"])
	fmt.Println(env_data_result["User"])
	fmt.Println(env_data_result["Password"])

	count := env_data_result["Interval"].(int)

	time.Sleep(time.Duration(count) * time.Second)

}
