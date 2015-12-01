package goping

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Ping struct is the Go representation of the data from a JSON ping.
type Ping struct {
	Origin           string `json:"origin"`
	NameLookupTimeMs int    `json:"name_lookup_time_ms"`
	ConnectTimeMs    int    `json:"connect_time_ms"`
	TransferTimeMs   int    `json:"transfer_time_ms"`
	TotalTimeMs      int    `json:"total_time_ms"`
	CreatedAt        string `json:"created_at"`
	Status           int    `json:"status"`
}

// Decodes the Ping array from the JSON data found in the body of the
// `http.Request` object.
func ReadPings(r *http.Request) ([]Ping, error) {
	var pings []Ping
	err := json.NewDecoder(r.Body).Decode(&pings)

	if err != nil {
		return nil, errors.New("Ill formed JSON ping.")
	}

	return pings, nil
}
