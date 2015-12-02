package tsdb

import utils_json "github.com/aseure/goping/utils/json"

// Interface for any connector to an arbitrary TSDB.
type TSDBConnector interface {
	// Add new datapoints to the TSDB.
	AddPings(pings []utils_json.Ping)
	// Retrieve a slice of the average transfer time of `origin` (in ms)
	// aggregated by hours from the oldest timestamp until now.
	GetAveragePerHour(origin string) []float64
	// Retrieve a slice of string containing all origin values
	GetOrigins() []string
}
