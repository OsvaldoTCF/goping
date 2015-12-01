package tsdb

import (
	"time"

	utils_json "github.com/aseure/goping/utils/json"
)

// Interface for any connector to an arbitrary TSDB.
type TSDBConnector interface {
	// Add new datapoints to the TSDB.
	AddPings(pings []utils_json.Ping)
	// Retrieve a slice of the average transfer time of `origin` (in ms)
	// aggregated by hours within the time range provided.
	GetAveragePerHour(origin string, start time.Time, end time.Time) []utils_json.Average
}
