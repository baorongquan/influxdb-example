package query

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const DATABASE = "test"

func Query(c client.Client) {
	q := client.NewQuery("SELECT count(busy) FROM cpu_usage", DATABASE, "ns")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}
}
