package goping

import (
	"time"

	"github.com/influxdb/influxdb/client/v2"
)

// Connector to an InfluxDB instance to handle pings as timeseries
type InfluxConnector struct {
	c client.Client
}

// Instantiates a new InfluxConnector
func NewInfluxConnector() *InfluxConnector {
	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})

	if err != nil {
		return nil
	}

	return &InfluxConnector{
		c: influxClient,
	}
}

func (connector *InfluxConnector) Connect() {
}

func (connector *InfluxConnector) AddPing(p *Ping) {
}

func (connector *InfluxConnector) GetAveragePerHour(start time.Time, end time.Time) []int {
	return make([]int, 1)
}
