package main

import (
	"github.com/brq/influxdb-example/connect"
	"github.com/brq/influxdb-example/database"
	// "github.com/brq/influxdb-example/write"
	"github.com/brq/influxdb-example/query"
)

func main() {
	c := connect.NewHTTPClient()
	defer c.Close()

	database.CreateDatabase(c)
	// database.DropDatabase(c)

	// write.WriteData(c)
	// write.BatchWriteData(c)

	query.Query(c)
}
