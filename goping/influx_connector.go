package goping

import (
	"log"
	"time"

	"github.com/influxdb/influxdb/client/v2"
)

// Connector to an InfluxDB instance to handle pings as timeseries.
type InfluxConnector struct {
	c        client.Client
	database string
}

// Instantiates a new InfluxConnector and connects to the InfluxDB instance.
func NewInfluxConnector() *InfluxConnector {
	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})

	if err != nil {
		return nil
	}

	return &InfluxConnector{
		c:        influxClient,
		database: "goping",
	}
}

func (connector *InfluxConnector) AddPing(p *Ping) {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  connector.database,
		Precision: "s",
	})

	tags := map[string]string{
		"origin": p.Origin,
	}

	fields := map[string]interface{}{
		"origin":              p.Origin,
		"name_lookup_time_ms": p.NameLookupTimeMs,
		"connect_time_ms":     p.ConnectTimeMs,
		"transfer_time_ms":    p.TransferTimeMs,
		"total_time_ms":       p.TotalTimeMs,
		"status":              p.Status,
	}

	timestamp, err := time.Parse("2006-02-01 15:04:05 MST", p.CreatedAt)
	if err != nil {
		log.Println(err)
	}

	point, err := client.NewPoint(
		"ping",
		tags,
		fields,
		timestamp,
	)

	bp.AddPoint(point)
	if err = connector.c.Write(bp); err != nil {
		log.Println("Cannot add the datapoint to InfluxDB.")
	}
}

func (connector *InfluxConnector) GetAveragePerHour(start time.Time, end time.Time) []int {
	return make([]int, 1)
}

// Query wrapper for InfluxDB commands.
func (connector *InfluxConnector) query(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: connector.database,
	}

	if response, err := connector.c.Query(q); err != nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	}

	return res, nil
}
