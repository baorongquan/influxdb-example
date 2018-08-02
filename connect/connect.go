package connect

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const (
	ADDR     = "http://localhost:8086"
	DATABASE = "test"
)

func main() {
	httpClient := newHTTPClient()
	defer httpClient.Close()

	udpClient := newUDPClient()
	defer udpClient.Close()
}

func newHTTPClient() client.Client {
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

func newUDPClient() client.Client {
	c, err := client.NewUDPClient(client.UDPConfig{
		Addr: ADDR,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
