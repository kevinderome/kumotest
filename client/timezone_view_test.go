package main

import (
	"testing"
	"time"
)

func TestDisplayWelcomeDialog(t *testing.T) {
	Display_welcome_dialog()
}

func TestDisplayTimes(t *testing.T) {
	var times = []time.Time{time.Now(), time.Now()}
	Display_times(&times)
}