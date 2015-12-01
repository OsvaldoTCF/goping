package main

import (
	"log"
	"net/http"

	"github.com/aseure/goping/tsdb"
	utils_json "github.com/aseure/goping/utils/json"
	"github.com/gorilla/mux"
)

var connector tsdb.TSDBConnector

func main() {
	// Instantiates and configures the connection to the TSDB
	connector = tsdb.NewInfluxConnector()

	// Instantiates and configures the HTTP router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/1/pings", PingAdd).Methods("POST")
	router.HandleFunc("/api/1/pings/{origin}/hours", PingGetAvgPerHour).Methods("GET")
	router.HandleFunc("/", PingChartWebPage).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler of /api/1/pings POST requests to add new datapoints in the TSDB.
func PingAdd(w http.ResponseWriter, r *http.Request) {
	pings, err := utils_json.ReadPings(r)
	if err != nil {
		log.Println(err)
	}

	connector.AddPings(pings)
}

// Handler of /api/1/pings/{origin}/hours GET requests to retrieve the everage
// `transfer_time_ms` for a specific `origin`, aggregated by hours.
func PingGetAvgPerHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	origin := vars["origin"]
	averages := connector.GetAveragePerHour(origin)

	utils_json.WriteAverages(w, averages)
}

// Handler of / GET requests to display a chart of the `transfer_time_ms`
// aggregated by hours with a dropdown menu to change the `origin`.
func PingChartWebPage(w http.ResponseWriter, r *http.Request) {
}
