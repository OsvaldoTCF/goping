package utils_json

import (
	"encoding/json"
	"log"
	"net/http"
)

type Average struct {
	Time  string `json:"t"`
	Value int    `json:"a"`
}

func WriteAverages(w http.ResponseWriter, averages []Average) {
	js, err := json.Marshal(averages)

	if err != nil {
		log.Println("Cannot marshalize aggregated averages.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
