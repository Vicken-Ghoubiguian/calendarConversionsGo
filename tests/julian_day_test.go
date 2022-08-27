package main

import (
	"time"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	var current_julian_day calendar_conversions.JulianDay

	//
	current_julian_day.Initialize_julian_day_from_time(time.Now())

	//
	current_julian_day.Format()
}