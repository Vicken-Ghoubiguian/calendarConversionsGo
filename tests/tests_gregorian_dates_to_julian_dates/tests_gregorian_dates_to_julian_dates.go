package main

import (
	"time"
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	fmt.Println("====== October Revolution ======")

	//
	var november_7_1917_gregorian_to_julian calendar_conversions.Julian
	var november_7_1917_from_time_golang_to_julian calendar_conversions.Julian
	var november_7_1917_in_gregorian calendar_conversions.Gregorian

	//
	november_7_1917_in_gregorian.Initialize_gregorian_from_elements(1917, 11, 7)
	november_7_1917_in_time_golang := time.Date(1917, time.Month(11), 7, 0, 0, 0, 0, time.UTC)

	//
	november_7_1917_gregorian_to_julian.Initialize_julian_from_Gregorian(november_7_1917_in_gregorian)
	november_7_1917_from_time_golang_to_julian.Initialize_julian_from_time_golang(november_7_1917_in_time_golang)

	//
	fmt.Println("Gregorian to Julian : " + november_7_1917_gregorian_to_julian.Format())
	
	//
	fmt.Println("-------------------------------")

	//
	fmt.Println("Golang time to Julian : " + november_7_1917_from_time_golang_to_julian.Format())

	//
	fmt.Println("================================")

	//
	fmt.Println("====== Proclamation of the first french Republic ======")

	//
	var september_22_1792_gregorian_to_julian calendar_conversions.Julian
	var september_22_1792_from_time_golang_to_julian calendar_conversions.Julian
	var september_22_1792_in_gregorian calendar_conversions.Gregorian

	//
	september_22_1792_in_gregorian.Initialize_gregorian_from_elements(1792, 9, 22)
	september_22_1792_in_time_golang := time.Date(1792, time.Month(9), 22, 0, 0, 0, 0, time.UTC)

	//
	september_22_1792_gregorian_to_julian.Initialize_julian_from_Gregorian(september_22_1792_in_gregorian)
	september_22_1792_from_time_golang_to_julian.Initialize_julian_from_time_golang(september_22_1792_in_time_golang)

	//

	//
	fmt.Println("-------------------------------")

	//

	//
	fmt.Println("=======================================================")
}