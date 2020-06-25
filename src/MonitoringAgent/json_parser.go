package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://openapi.cloudn.co.kr/cdnservice/statisticsapi/statistics/totaldomain/transfer")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		//var data map[string] interface{} //json 데이터를 저장할 공간을 맵으로 선언
		fmt.Println(string(contents))
		data := make(map[string]interface{})
		json.Unmarshal([]byte(contents), &data)
		for k, j := range data {
			fmt.Println("ex2 :", k, j)
		}
	}
}
