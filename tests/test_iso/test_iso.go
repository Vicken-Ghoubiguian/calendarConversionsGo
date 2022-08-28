package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	var current_date_in_iso calendar_conversions.ISO_week_based_calendar_date

	//
	current_date_in_iso.Initialize_ISO_week_based_calendar_date(time.Now())

	//
	fmt.Println(current_date_in_iso.Format(true))

	//
	fmt.Println(current_date_in_iso.Format(false))
}