package main

import (
	"testing"
)

// test GET method to Kumojin server
func TestGetTimeZone(t *testing.T) {
	_, err := Get_server_timezone()
	if err != nil {
		t.Fatal(err)
	}
}