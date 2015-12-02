package tsdb

import utils_json "github.com/aseure/goping/utils/json"

// Interface for any connector to an arbitrary TSDB.
type TSDBConnector interface {
	// Add new datapoints to the TSDB.
	AddPings(pings []utils_json.Ping)
	// Retrieve a slice of the average transfer time of `origin` (in ms) for
	// the first recorded 24 hours.
	GetAveragePerHour(origin string) *utils_json.AvgCollection
	// Retrieve a slice of string containing all origin values.
	GetOrigins() []string
}
