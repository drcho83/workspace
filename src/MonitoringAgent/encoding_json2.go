//Json 형태로 GO 받기
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"reflect"
	"strconv"

	"github.com/tidwall/gjson"
)

type BaseStruct struct {
	mainServerStatus string
	etc1             int64
	etc2             int64
	etc3             int64
	serverID         int64
	userCount        int64
	loginCount       int64
}

func main() {
	//마샬링: 논리적 구조를 바이트로 변경
	/*
		doc := `
			{
				"serverList":[
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":379188784,"etc2":5346768814,"etc3":4273911029521},"serverId":1101,"userCount":64,"loginCount":44},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1102,"userCount":53,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1103,"userCount":47,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1201,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1202,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1203,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1204,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":1205,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":364890109,"etc2":4943670966,"etc3":3916386837568},"serverId":2101,"userCount":47,"loginCount":41},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2102,"userCount":40,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2103,"userCount":36,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2201,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2202,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2203,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2204,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":2205,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":89152150,"etc2":1234123005,"etc3":3602676150776},"serverId":6101,"userCount":28,"loginCount":16},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6102,"userCount":28,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6103,"userCount":26,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6201,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6202,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6203,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6204,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":6205,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":617015801,"etc2":8668528080,"etc3":10524498266118},"serverId":7101,"userCount":86,"loginCount":67},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7102,"userCount":84,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7103,"userCount":90,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7201,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7202,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7203,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7204,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":7205,"userCount":0,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":6339770,"etc2":140429055,"etc3":8004762621173},"serverId":8101,"userCount":44,"loginCount":31},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8102,"userCount":45,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8103,"userCount":40,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8201,"userCount":3,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8202,"userCount":3,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8203,"userCount":3,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8204,"userCount":3,"loginCount":0},
				{"logicalServerStatus":{"mainServerStatus":true},"etc":{"etc1":0,"etc2":0,"etc3":0},"serverId":8205,"userCount":3,"loginCount":0}
				]
			}
			`
	*/
	var doc string
	response, err := http.Get("http://localhost/getstatus.json")
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

		//fmt.Println(string(contents))
		doc = string(contents)
		//fmt.Println(doc)

	}
	//declere valiable
	var MainServerStatus string
	var Etc1 string
	var Etc2 string
	var Etc3 string
	var ServerID string
	var UserCount string
	var LoginCount string

	data := make(map[string]interface{})

	//convert b, gjson.result to int64

	a := gjson.Get(doc, "serverList.#")
	stringB := a.String()
	b, err := strconv.Atoi(stringB)
	if err == nil {
		fmt.Println(b)
	}

	for i := 0; i <= 40; i++ {
		MainServerStatus = "serverList." + strconv.Itoa(i) + ".logicalServerStatus.mainServerStatus"
		Etc1 = "serverList." + strconv.Itoa(i) + ".etc.etc1"
		Etc2 = "serverList." + strconv.Itoa(i) + ".etc.etc2"
		Etc3 = "serverList." + strconv.Itoa(i) + ".etc.etc3"
		ServerID = "serverList." + strconv.Itoa(i) + ".serverId"
		UserCount = "serverList." + strconv.Itoa(i) + ".userCount"
		LoginCount = "serverList." + strconv.Itoa(i) + ".loginCount"

		//Get Json
		ServerID := gjson.Get(doc, ServerID)                 //gjson Type
		MainServerStatus := gjson.Get(doc, MainServerStatus) //gjson Type
		Etc1 := gjson.Get(doc, Etc1)                         //gjson Type
		Etc2 := gjson.Get(doc, Etc2)                         //gjson Type
		Etc3 := gjson.Get(doc, Etc3)                         //gjson Type
		UserCount := gjson.Get(doc, UserCount)               //gjson Type
		LoginCount := gjson.Get(doc, LoginCount)             //gjson Type

		//fmt.Println(reflect.TypeOf(ServerID))
		//fmt.Println(reflect.TypeOf(ServerID.Int()))
		//fmt.Println(ServerID.String())
		//fmt.Println(MainServerStatus.String())
		if ServerID.Int() != 0 {
			data[strconv.Itoa(i)] = BaseStruct{
				serverID:         ServerID.Int(),
				mainServerStatus: MainServerStatus.String(),
				etc1:             Etc1.Int(),
				etc2:             Etc2.Int(),
				etc3:             Etc3.Int(),
				userCount:        UserCount.Int(),
				loginCount:       LoginCount.Int()}
		}
		//fmt.Println("BaseStruct: ", data[strconv.Itoa(i)])
		//data[strconv.Itoa(i)] = value
	}
	for _, j := range data {
		fmt.Println(j)
	}
}
