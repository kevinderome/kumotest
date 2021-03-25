package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/timezone", TimeZoneHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}