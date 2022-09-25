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
	var current_julian_date calendar_conversions.Julian

	//
	current_julian_date.Initialize_julian_from_time_golang(time.Now())

	//
	fmt.Println(current_julian_date.Format())

	//
	current_julian_day.Initialize_julian_day_from_julian_date(current_julian_date)

	//
	fmt.Println(current_julian_day.Format())
}