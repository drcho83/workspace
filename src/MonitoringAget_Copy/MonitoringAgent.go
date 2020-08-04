package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/tidwall/gjson"
	//"reflect"
	client "github.com/influxdata/influxdb1-client/v2"
)

type BaseStruct struct {
	MainServerStatus string
	Etc1             int64
	Etc2             int64
	Etc3             int64
	Etc4             int64
	ServerID         int64
	UserCount        int64
	LoginCount       int64
}

const (
	MyDB     = "test1"
	username = "user"
	password = "pw"
)

func main() {

	wg := new(sync.WaitGroup)

	defile_map := map[string]string{
		"L2M_STAT2": "http://localhost/l2m.json",
		"H2_STAT2":  "http://localhost/h2.json",
		"LM_STAT2":  "http://localhost/lm.json",
	}

	for k, v := range defile_map {
		fmt.Println("GetInfo: ", k, v)
		wg.Add(1)

		go func(gameid string, url string) {

			var doc string
			response, err := http.Get(url)
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

			//declare valiable
			var mainServerStatus string
			var etc1 string
			var etc2 string
			var etc3 string
			var etc4 string
			var serverID string
			var userCount string
			var loginCount string

			data := make(map[string]BaseStruct)

			//convert b, gjson.result to int64

			//Json 파일의 Line Count를 가져온다.
			a := gjson.Get(doc, "serverList.#")
			stringB := a.String()
			linecount, err := strconv.Atoi(stringB)
			if err == nil {
				fmt.Println(linecount)
			}

			//define influxdb connection
			c, err := client.NewHTTPClient(client.HTTPConfig{
				Addr:     "http://192.168.65.133:8086",
				Username: "admin",
				Password: "Pa$$w0rd",
			})
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()

			// end influxdb conntion

			for i := 0; i <= linecount; i++ {
				mainServerStatus = "serverList." + strconv.Itoa(i) + ".logicalServerStatus.mainServerStatus"
				etc1 = "serverList." + strconv.Itoa(i) + ".etc.etc1"
				etc2 = "serverList." + strconv.Itoa(i) + ".etc.etc2"
				etc3 = "serverList." + strconv.Itoa(i) + ".etc.etc3"
				etc4 = "serverList." + strconv.Itoa(i) + ".etc.etc4"
				serverID = "serverList." + strconv.Itoa(i) + ".serverId"
				userCount = "serverList." + strconv.Itoa(i) + ".userCount"
				loginCount = "serverList." + strconv.Itoa(i) + ".loginCount"

				//Get Json
				serverID := gjson.Get(doc, serverID)                 //gjson Type
				mainServerStatus := gjson.Get(doc, mainServerStatus) //gjson Type
				etc1 := gjson.Get(doc, etc1)                         //gjson Type
				etc2 := gjson.Get(doc, etc2)                         //gjson Type
				etc3 := gjson.Get(doc, etc3)                         //gjson Type
				etc4 := gjson.Get(doc, etc4)                         //gjson Type
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
						Etc4:             etc4.Int(),
						UserCount:        userCount.Int(),
						LoginCount:       loginCount.Int()}
				}

				/*
					if serverID.Int() != 0 {
						fmt.Println(data[strconv.Itoa(i)].ServerID, data[strconv.Itoa(i)].MainServerStatus, data[strconv.Itoa(i)].Etc1, data[strconv.Itoa(i)].Etc2, data[strconv.Itoa(i)].Etc3, data[strconv.Itoa(i)].UserCount, data[strconv.Itoa(i)].LoginCount)
				}*/

				//create batch points
				bp, err := client.NewBatchPoints(client.BatchPointsConfig{
					Database:  "test1",
					Precision: "ms",
				})
				if err != nil {
					log.Fatal(err)
				}
				// Create a point and add to batch
				tags := map[string]string{"ServerID": strconv.FormatInt(data[strconv.Itoa(i)].ServerID, 10)}
				//var test1 int64
				//var test2 string
				//test1 = data[strconv.Itoa(i)].ServerID
				//test2 = strconv.FormatInt(test1, 10)
				//fmt.Println(test1, test2)

				//tags := map[string]int{"ServerID": data.ServerID}
				fields := map[string]interface{}{
					"MainServerStatus": data[strconv.Itoa(i)].MainServerStatus,
					"Etc1":             data[strconv.Itoa(i)].Etc1,
					"Etc2":             data[strconv.Itoa(i)].Etc2,
					"Etc3":             data[strconv.Itoa(i)].Etc3,
					"Etc4":             data[strconv.Itoa(i)].Etc4,
					"UserCount":        data[strconv.Itoa(i)].UserCount,
					"LoginCount":       data[strconv.Itoa(i)].LoginCount,
				}

				pt, err := client.NewPoint(gameid, tags, fields, time.Now())
				if err != nil {
					log.Fatal(err)
				}

				bp.AddPoint(pt)

				// Write the batch
				if err := c.Write(bp); err != nil {
					log.Fatal(err)
				}

				// Close client resources
				if err := c.Close(); err != nil {
					log.Fatal(err)
				}
			}
			wg.Done()
		}(k, v)
	}

	wg.Wait()
	fmt.Println("insert Complete!")

}
