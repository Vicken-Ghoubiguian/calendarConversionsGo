package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	var current_julian_day calendar_conversions.JulianDay

	//
	current_julian_day.Initialize_julian_day_from_time(time.Now())

	//
	fmt.Println(current_julian_day.Format())
}