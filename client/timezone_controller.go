package main

import (
	"time"
	"fmt"
)

// represent the server timezone
var server_timezone string = ""

// convert timezone name to time.Time valid.
func timeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func Welcome_dialog_controller() {
	Display_welcome_dialog()
}

// Controller for timezone, this function get the server timezone and forward to display_times view.
func Timezone_controller() {
	if server_timezone == "" {
		zone, err := Get_server_timezone()
		if err != nil {
			fmt.Println("Impossible to call /timezone endpoint")
		} else {
			server_timezone = zone
		}
	}
	server_time, err := timeIn(time.Now(), server_timezone)
		var times = make([]time.Time, 0)
	if err == nil {
		times = append(times, server_time)
	}
	times = append(times, time.Now())
	Display_times(&times)
}