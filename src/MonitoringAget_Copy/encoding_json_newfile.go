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
	MainServerStatus string
	Etc1             int64
	Etc2             int64
	Etc3             int64
	ServerID         int64
	UserCount        int64
	LoginCount       int64
}

func main() {

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
	var mainServerStatus string
	var etc1 string
	var etc2 string
	var etc3 string
	var serverID string
	var userCount string
	var loginCount string

	data := make(map[string]BaseStruct)

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
		//j, _ := json.Marshal(data)
		//json.Unmarshal(j, &data)
		//fmt.Println(&data)
		//fmt.Println(data)
		if serverID.Int() != 0 {
			fmt.Println(data[strconv.Itoa(i)].ServerID, data[strconv.Itoa(i)].MainServerStatus, data[strconv.Itoa(i)].Etc1, data[strconv.Itoa(i)].Etc2, data[strconv.Itoa(i)].Etc3, data[strconv.Itoa(i)].UserCount, data[strconv.Itoa(i)].LoginCount)
		}
	}
	/*
		for k, j := range data {
			fmt.Println(k, j)
		}
	*/
}
