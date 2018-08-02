package write

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

const (
	DATABASE = "test"
)

func WriteData(c client.Client) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database:  DATABASE,
	})

	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	// Create a point and add to batch
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}
