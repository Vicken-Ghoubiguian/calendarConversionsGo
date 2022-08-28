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
	current_julian_date.Initialize_julian(time.Now())

	//
	fmt.Println(current_julian_date.Format())
}