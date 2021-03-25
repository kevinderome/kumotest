package main

import (
	"fmt"
	"time"
)

// Display the welcome dialog to whish welcome to the users.
func Display_welcome_dialog() {
	fmt.Println("⡅⠥⠍⠕⠚⠊⠝        Kumojin")
	fmt.Println("Welcome to you!")
	fmt.Println("I'll give you the time of your home as well as the time of our favorite server ;)")
	fmt.Println("----------")
}

// View for display times
// @times is a time.Time slice
// all items are displayed
func Display_times(times *[]time.Time) {
	i := 0
	for i < len(*times) {
		fmt.Println((*times)[i].Location(), (*times)[i].Format("15:04"))
		fmt.Println("**********")
		i = i + 1
	}
}