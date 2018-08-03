package query

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const DATABASE = "test"

func Query(c client.Client) {
	q := client.NewQuery("SELECT * FROM cpu_usage Limit 5", DATABASE, "ns")
	response, err := c.Query(q)
	if err == nil {
		if response.Error() != nil {
			log.Println(" err :", response.Error())
			return
		}
	} else {
		log.Fatalln(err)
		return
	}

	res := response.Results
	for _, r := range res {
		for _, s := range r.Series {
			log.Println("tag :", s.Tags)
			log.Println("columns:", s.Columns)
			for _, v := range s.Values {
				log.Println("fields :", v)
			}
		}
	}
}
