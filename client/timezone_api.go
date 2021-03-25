package main

import (
	"log"
	"net/http"
	"encoding/json"
)

// This structure represent a TimeZone object
type TimeZone struct {
	TimeZone string `json:"TimeZone"`
}
// GET the server TimeZone
func Get_server_timezone() (timezone string, err error) {
	resp, err := http.Get("http://localhost:8001/timezone")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var result TimeZone
	json.NewDecoder(resp.Body).Decode(&result)
	return result.TimeZone, nil
}