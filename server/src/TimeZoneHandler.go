package main

import (
	"net/http"
	"encoding/json"
	"log"
)

// This constant defines the desired time zone by default.
const defaultTimeZone string = "Asia/Tokyo"

// This structure represent a TimeZone object
type TimeZone struct {
	TimeZone string `json:"TimeZone"`
}

// This function allows to display an error as well as to return an HTTP error.
func repportErrorHandler(w *http.ResponseWriter, err string, code int) {
	log.Println(err)
	http.Error(*w, err, code)
}

// This function corresponds to the "/timezone" route.
// Only the GET method is supported.
// In case of success the method sends back to the client a json containing the timezone of the server.
// else if fail, server send back HTTP code appropriate.
func TimeZoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		timezone := TimeZone{defaultTimeZone}

		if err := json.NewEncoder(w).Encode(timezone); err != nil {
			repportErrorHandler(&w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		repportErrorHandler(&w, "Bad method Request - This route only accepts GET requests", http.StatusBadRequest)
	}
}