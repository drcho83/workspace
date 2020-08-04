package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	//"math/rand"

	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	MyDB     = "test1"
	username = "user"
	password = "pw"
)

func main() {

	wg := new(sync.WaitGroup)

	arr1 := [3]string{"LM", "L2M", "H2"}

	for _, v := range arr1 {
		fmt.Println("ex1: ", v)
		wg.Add(1)

		go func(table string) {
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
			for i := 1; i <= 9; i++ {

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
					"user":   70.1 +float64(i),
				}

				pt, err := client.NewPoint(table, tags, fields, time.Now())
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
				//time.Sleep(800 * time.Millisecond)
			}
			wg.Done()
		}(v)
	}

	wg.Wait()
	fmt.Println("insert Complete!")

}
