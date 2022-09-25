package main

import (
	"fmt"
	"github.com/Vicken-Ghoubiguian/calendar_conversions"
)

func main() {

	//
	fmt.Println("====== October Revolution ======")

	//
	var november_7_1917_gregorian_to_julian calendar_conversions.Julian
	var november_7_1917_in_gregorian calendar_conversions.Gregorian

	//
	november_7_1917_in_gregorian.Initialize_gregorian_from_elements(1917, 11, 7, 0, 0, 0, 0)

	//
	november_7_1917_gregorian_to_julian.Initialize_julian_from_Gregorian(november_7_1917_in_gregorian)

	//
	fmt.Println("Julian : " + november_7_1917_gregorian_to_julian.Format())

	//
	fmt.Println("================================")

	//
	fmt.Println("====== Proclamation of the first french Republic ======")

	//
	var september_22_1792_gregorian_to_julian calendar_conversions.Julian
	var september_22_1792_in_gregorian calendar_conversions.Gregorian

	//
	september_22_1792_in_gregorian.Initialize_gregorian_from_elements(1792, 9, 22, 0, 0, 0, 0)

	//
	september_22_1792_gregorian_to_julian.Initialize_julian_from_Gregorian(september_22_1792_in_gregorian)

	//
	fmt.Println("Julian : " + september_22_1792_gregorian_to_julian.Format())

	//
	fmt.Println("================================")
}