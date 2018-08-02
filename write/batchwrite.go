package write

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"math/rand"
	"time"
)

const size = 1000

func BatchWriteData(c client.Client) {
	// Create a new point batch
	rand.Seed(45)
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "us",
		Database:  DATABASE,
	})

	for i := 0; i < size; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu":     "cpu-total",
			"host":    fmt.Sprintf("host%d", rand.Intn(1000)),
			"regions": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		// Create a point and add to batch
		pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
		if err != nil {
			log.Println("Error: ", err.Error())
			continue
		}
		bp.AddPoint(pt)
	}
	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal("Error: ", err.Error())
	}
}
