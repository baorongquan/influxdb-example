package database

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

func CreateDatabase(c client.Client) {
	q := client.NewQuery("CREATE DATABASE test", "", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}
}

func DropDatabase(c client.Client) {
	q := client.NewQuery("DROP DATABASE test", "", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}
}
