package goping

import "time"

// Interface for any connector to an arbitrary TSDB.
type TSDBConnector interface {
	// Add new datapoints to the TSDB.
	AddPings(pings []Ping)
	// Retrieve a slice of the average transfer time (in ms) aggregated by hours
	// between the time range provided.
	GetAveragePerHour(start time.Time, end time.Time) []int
}
