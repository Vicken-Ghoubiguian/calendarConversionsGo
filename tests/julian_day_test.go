package main

import (
	"time"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	var test calendar_conversions.JulianDay

	//
	test.Initialize_julian_day_from_time(time.Now())

	//
	test.Format()
}