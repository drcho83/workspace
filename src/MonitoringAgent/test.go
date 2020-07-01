package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Persion struct {
	Id       int
	Name     string
	Address  string
	Email    string
	School   string
	City     string
	Company  string
	Age      int
	Sex      string
	Proviece string
	Com      string
	PostTo   string
	Buys     string
	Hos      string
}

func main() {
	StructToMapViaJson()
	//StructToMapViaReflect()
}

func StructToMapViaJson() {
	m := make(map[string]interface{})
	t := time.Now()
	person := Persion{
		Id:       98439,
		Name:     "zhaondifnei",
		Address:  "Great Sandy Land.",
		Email:    "dashdisnin@126.com",
		School:   "Guangzhou No. 15 Middle School.",
		City:     "zhongguoguanzhou",
		Company:  "sndifneinsifnienisn",
		Age:      23,
		Sex:      "F",
		Proviece: "jianxi",
		Com:      "Lamborghini, Guangzhou",
		PostTo:   "Blue Whale XXXXXXXX",
		Buys:     "shensinfienisnfieni",
		Hos:      "zhonsndifneisnidnfie",
	}
	j, _ := json.Marshal(person)
	json.Unmarshal(j, &m)
	fmt.Println(m)
	fmt.Println(time.Now().Sub(t))
}
