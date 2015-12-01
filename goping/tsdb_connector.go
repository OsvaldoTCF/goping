package goping

import "time"

// Interface for any connector to an arbitrary TSDB.
type TSDBConnector interface {
	// Initializes the connection with the TSDB instance.
	Connect()
	// Add a new datapoint to the TSDB.
	AddPing(p *Ping)
	// Retrieve a slice of the average transfer time (in ms) aggregated by hours
	// between the time range provided.
	GetAveragePerHour(start time.Time, end time.Time) []int
}
