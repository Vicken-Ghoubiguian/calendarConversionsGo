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
	var current_gregorian_date calendar_conversions.Gregorian

	//
	current_gregorian_date.Initialize_gregorian_from_time(time.Now())

	fmt.Println(current_gregorian_date.Format())

	//
	current_julian_day.Initialize_julian_day_from_gregorian_date(current_gregorian_date)

	//
	fmt.Println(current_julian_day.Format())
}