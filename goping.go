package main

import (
	"fmt"
	"log"
	"net/http"

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
	router.HandleFunc("/api/1/pings/{origin}/hours", HandlerGetAvgPerHour).Methods("GET")
	router.HandleFunc("/", HandlerChartWebPage).Methods("GET")

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
func HandlerGetAvgPerHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	origin := vars["origin"]
	avgCollection := connector.GetAveragePerHour(origin)

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
