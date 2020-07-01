//Json 형태로 GO 받기
package main

import (
	"encoding/json"
	"fmt"

	//"io/ioutil"
	//"net/http"
	//"os"

	//"reflect"
	"strconv"

	"github.com/tidwall/gjson"
)

type BaseStruct struct {
	MainServerStatus string
	Etc1             int64
	Etc2             int64
	Etc3             int64
	ServerID         int64
	UserCount        int64
	LoginCount       int64
}

func main() {
	//마샬링: 논리적 구조를 바이트로 변경

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
	/*
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
	*/
	//declere valiable
	var mainServerStatus string
	var etc1 string
	var etc2 string
	var etc3 string
	var serverID string
	var userCount string
	var loginCount string

	data := make(map[string]interface{})

	//convert b, gjson.result to int64

	a := gjson.Get(doc, "serverList.#")
	stringB := a.String()
	b, err := strconv.Atoi(stringB)
	if err == nil {
		fmt.Println(b)
	}

	for i := 0; i <= b; i++ {
		mainServerStatus = "serverList." + strconv.Itoa(i) + ".logicalServerStatus.mainServerStatus"
		etc1 = "serverList." + strconv.Itoa(i) + ".etc.etc1"
		etc2 = "serverList." + strconv.Itoa(i) + ".etc.etc2"
		etc3 = "serverList." + strconv.Itoa(i) + ".etc.etc3"
		serverID = "serverList." + strconv.Itoa(i) + ".serverId"
		userCount = "serverList." + strconv.Itoa(i) + ".userCount"
		loginCount = "serverList." + strconv.Itoa(i) + ".loginCount"

		//Get Json
		serverID := gjson.Get(doc, serverID)                 //gjson Type
		mainServerStatus := gjson.Get(doc, mainServerStatus) //gjson Type
		etc1 := gjson.Get(doc, etc1)                         //gjson Type
		etc2 := gjson.Get(doc, etc2)                         //gjson Type
		etc3 := gjson.Get(doc, etc3)                         //gjson Type
		userCount := gjson.Get(doc, userCount)               //gjson Type
		loginCount := gjson.Get(doc, loginCount)             //gjson Type

		//fmt.Println(reflect.TypeOf(ServerID))
		//fmt.Println(reflect.TypeOf(ServerID.Int()))
		//fmt.Println(ServerID.String())
		//fmt.Println(MainServerStatus.String())
		if serverID.Int() != 0 {
			data[strconv.Itoa(i)] = BaseStruct{
				ServerID:         serverID.Int(),
				MainServerStatus: mainServerStatus.String(),
				Etc1:             etc1.Int(),
				Etc2:             etc2.Int(),
				Etc3:             etc3.Int(),
				UserCount:        userCount.Int(),
				LoginCount:       loginCount.Int()}
		}
		//fmt.Println("BaseStruct: ", data[strconv.Itoa(i)])
		//data[strconv.Itoa(i)] = value
		//fmt.Println(data)
		j, _ := json.Marshal(data)
		json.Unmarshal(j, &data)
		//fmt.Println(&data)
		//fmt.Println(data)
	}

	for k, j := range data {
		fmt.Println(k, j)
	}

}
