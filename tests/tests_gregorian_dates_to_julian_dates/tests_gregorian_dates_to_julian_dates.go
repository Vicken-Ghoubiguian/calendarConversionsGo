package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	fmt.Println("===============================")
	var november_7_1917_in_julian calendar_conversions.Julian
	var november_7_1917_in_gregorian calendar_conversions.Gregorian
	november_7_1917_in_time_golang := time.Date(1917, time.Month(11), 7, 0, 0, 0, 0, time.UTC)


	fmt.Println("===============================")
}