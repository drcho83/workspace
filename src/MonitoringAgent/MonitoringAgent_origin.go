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

	//arr1 := [3]string{"LM", "L2M", "H2"}

	defile_map := map[string]string{
		"LMS": "http://localhost/getstatus1.json",
		//	"L2M": "http://localhost/getstatus2.json",
		//	"H2":  "http://localhost/getstatus3.json",
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
			line_count, err := strconv.Atoi(stringB)
			if err == nil {
				fmt.Println(line_count)
			}

			for i := 0; i <= line_count; i++ {
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

			//influxDB Insert!!!
			c, err := client.NewHTTPClient(client.HTTPConfig{
				Addr:     "http://192.168.65.132:8086",
				Username: "admin",
				Password: "Pa$$w0rd",
			})
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()

			// Create a new point batch
			for i := 0; i <= 10; i++ {
				bp, err := client.NewBatchPoints(client.BatchPointsConfig{
					Database:  "test1",
					Precision: "ms",
				})
				if err != nil {
					log.Fatal(err)
				}

				// Create a point and add to batch
				tags := map[string]string{"cpu": "cpu-total"}
				fields := map[string]interface{}{
					"idle":   10.1,
					"system": 53.3,
					"user":   80.1 + float64(i),
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
