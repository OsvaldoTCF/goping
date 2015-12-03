package tsdb

import (
	"time"

	utils_json "github.com/aseure/goping/utils/json"
)

// Interface for any connector to an arbitrary TSDB.
type TSDBConnector interface {
	// Add new datapoints to the TSDB.
	AddPings(pings []utils_json.Ping)

	// Retrieve a slice of the average transfer time of `origin` (in ms) for
	// the first recorded day.
	GetAveragePerHour(origin string) *utils_json.AvgCollection

	// Retrieve a slice of the average transfer time of `origin` (in ms) for
	// the current day.
	GetAveragePerHourNow(origin string) *utils_json.AvgCollection

	// Generic method to retrieve any array of averages.
	// For instance, if we need to retrieve averages per hour of the last 24
	// hours, the parameters must be set to:
	//
	//   - start: time.Now().AddDate(0, 0, -1)
	//   - step: time.Hour
	//   - count: 24
	//
	GetAverages(origin string, start time.Time, step time.Duration, count int) *utils_json.AvgCollection

	// Retrieve a slice of string containing all origin values.
	GetOrigins() []string
}
