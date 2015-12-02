package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aseure/goping/tsdb"
	utils_json "github.com/aseure/goping/utils/json"
	"github.com/aseure/goping/webview"
	"github.com/gorilla/mux"
)

var connector tsdb.TSDBConnector

func main() {
	// Instantiates and configures the connection to the TSDB
	connector = tsdb.NewInfluxConnector()

	// Instantiates and configures the HTTP router
	router := mux.NewRouter().StrictSlash(true)

	// API endpoints
	router.HandleFunc("/api/1/pings", HandlerAdd).Methods("POST")
	router.HandleFunc("/api/1/pings/{origin}/hours", HandlerGetAvgPerHourV1).Methods("GET")
	router.HandleFunc("/", HandlerChartWebPage).Methods("GET")
	router.HandleFunc("/api/2/pings/{origin}/{time}/{way}", HandlerGetAvgPerHourV2).Methods("GET")

	// Static ressources
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("webview/css/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("webview/js/"))))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("webview/fonts/"))))

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler of /api/1/pings POST requests to add new datapoints in the TSDB.
func HandlerAdd(w http.ResponseWriter, r *http.Request) {
	pings, err := utils_json.ReadPings(r)
	if err != nil {
		log.Println(err)
	}

	connector.AddPings(pings)
}

// Handler of /api/1/pings/{origin}/hours GET requests to retrieve the average
// `transfer_time_ms` for a specific `origin`, aggregated by hours.
func HandlerGetAvgPerHourV1(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	origin := vars["origin"]
	avgCollection := connector.GetAveragePerHour(origin)

	utils_json.WriteAverages(w, avgCollection)
}

// Handler of /api/2/pings/{origin}/{time} GET requests to retrieve the average
// `transfer_time_ms` for a specific `origin`, aggregated by hours. The first
// date is {time + 1 day} if way is set to "next" or {time - 1 day} if way is
// set to "prev".
func HandlerGetAvgPerHourV2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	origin := vars["origin"]
	start, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", vars["time"])
	way := vars["way"]
	avgCollection := utils_json.NewAvgCollection(0)

	if err == nil {
		if way == "prev" {
			start = start.AddDate(0, 0, -1)
			avgCollection = connector.GetAverages(origin, start, time.Hour, 24)
		} else if way == "next" {
			start = start.AddDate(0, 0, +1)
			avgCollection = connector.GetAverages(origin, start, time.Hour, 24)
		}
	}

	utils_json.WriteAverages(w, avgCollection)
}

// Handler of / GET requests to display a chart of the `transfer_time_ms`
// aggregated by hours with a dropdown menu to change the `origin`.
func HandlerChartWebPage(w http.ResponseWriter, r *http.Request) {
	p, err := webview.LoadPage("webview/index.html")
	if err != nil {
		log.Println("Cannot find `webview/index.html`.")
		writePageNotFound(w)
		return
	}

	origins := connector.GetOrigins()

	if err = p.WritePage(w, origins); err != nil {
		log.Println("Cannot render `webview/index.html`.")
		writeInternalServerError(w)
	}
}

// Write a 404 error on the ResponseWriter
func writePageNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 page not found")
}

// Write a 500 error on the ResponseWriter
func writeInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "500 internal server error")
}
