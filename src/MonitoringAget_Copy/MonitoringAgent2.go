package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/influxdata/influxdb1-client/v2"
)

func writePoints(clnt client.Client) {
	sampleSize := 1000

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "systemstats",
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sampleSize; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu":    "cpu-total",
			"host":   fmt.Sprintf("host%d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		pt, err := client.NewPoint(
			"cpu_usage",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}
func main() {
}
