package main

import (
	"log"
	"time"

	"github.com/influxdata/influxdb1-client/v2"
)

const (
	MyDB     = "test1"
	username = "user"
	password = "pw"
)

//func insertData(ch chan bool) {
func test() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.65.128:8086",
		Username: "admin",
		Password: "Pa$$w0rd",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.99,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
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

func main() {
	test()
}
