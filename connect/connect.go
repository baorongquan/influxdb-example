package connect

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const (
	ADDR     = "http://localhost:8086"
	DATABASE = "test"
)

func NewHTTPClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: ADDR,
		// Username: "username",
		// Password: "password",
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func NewUDPClient() client.Client {
	c, err := client.NewUDPClient(client.UDPConfig{
		Addr: ADDR,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
