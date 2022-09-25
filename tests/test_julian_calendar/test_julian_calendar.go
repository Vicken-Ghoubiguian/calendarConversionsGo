package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	var current_julian_date calendar_conversions.Julian

	//
	var current_gregorian_date calendar_conversions.Gregorian

	//
	current_gregorian_date.Initialize_gregorian_from_time(time.Now())

	//
	current_julian_date.Initialize_julian_from_Gregorian(current_gregorian_date)

	//
	fmt.Println(current_julian_date.Format())
}