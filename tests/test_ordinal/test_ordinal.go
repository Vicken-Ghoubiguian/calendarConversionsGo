package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	var current_date_in_ordinal calendar_conversions.Ordinal_date

	//
	current_date_in_ordinal.Initialize_ISO_week_based_calendar_date(time.Now())

	//
	fmt.Println(current_date_in_ordinal.Format())
}