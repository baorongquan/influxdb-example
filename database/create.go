package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const (
	ADDR     = "http://localhost:8086"
	DATABASE = "test"
)

func main() {
	CreateDatabase()
	DropDatabase()
}

func CreateDatabase() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: ADDR,
		// Username: "username",
		// Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	q := client.NewQuery("CREATE DATABASE test", "", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}
}

func DropDatabase() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: ADDR,
		// Username: "username",
		// Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	q := client.NewQuery("DROP DATABASE test", "", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}
}
