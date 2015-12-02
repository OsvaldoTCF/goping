package utils_json

import "time"

// AvgCollection struct is the Go representation of a collection of averages
// values and timestamps.
type AvgCollection struct {
	Averages []float64 `json:"averages"`
	Times    []string  `json:"times"`
}

// Instantiates a new AvgCollection.
func NewAvgCollection(start string, step time.Duration, count int) *AvgCollection {
	return &AvgCollection{
		Averages: make([]float64, count),
		Times:    make([]string, count),
	}
}
