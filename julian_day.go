package calendar_conversions

//
import (
	"time"
	"fmt"
	"math"
)

//
type JulianDay struct {

	//
	countOfDaysSinceJulianPeriod float64
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_time(dt time.Time) {

	//
	Y := float64(dt.Year())
	M := float64(dt.Month())
	D := float64(dt.Day())

	//
	A := math.Floor(Y/100)
	B := math.Floor(A/4)
	C := math.Floor(2 - A + B)
	E := math.Floor(365.25 * (Y + 4716))
	F := math.Floor(30.6001 * (M + 1))

	//
	/*H := dt.Hour() / 24.0
    MIN := math.Round((dt.Minute() / (24.0 * 60)), 5)
    SEC := math.Round((dt.Second() / (24.0 * 60 * 60)), 5)
    MILLISEC := dt.Microseconds() / (24.0 * 60 * 60 * 1000)*/

	//
	current_julian_day.countOfDaysSinceJulianPeriod = (C + D + E + F - 1524.5) // + H + MIN + SEC + MILLISEC
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_gregorian_date(gregorian_date Gregorian) {

	//
	Y := float64(gregorian_date.Get_year())
	M := float64(gregorian_date.Get_monthNumber())
	D := float64(gregorian_date.Get_day())

	//
	A := math.Floor(Y/100)
	B := math.Floor(A/4)
	C := math.Floor(2 - A + B)
	E := math.Floor(365.25 * (Y + 4716))
	F := math.Floor(30.6001 * (M + 1))

	//
	/*H := dt.Hour() / 24.0
    MIN := math.Round((dt.Minute() / (24.0 * 60)), 5)
    SEC := math.Round((dt.Second() / (24.0 * 60 * 60)), 5)
    MILLISEC := dt.Microseconds() / (24.0 * 60 * 60 * 1000)*/

	//
	current_julian_day.countOfDaysSinceJulianPeriod = (C + D + E + F - 1524.5) // + H + MIN + SEC + MILLISEC
}

func (current_julian_day *JulianDay) Initialize_julian_day_from_julian_date(julian_date Julian) {

	
}

//
func (current_julian_day *JulianDay) Determine_gregorian_date() Gregorian {

	//
	var correspondingGregorianDate Gregorian

	//
	/*Q := current_julian_day.countOfDaysSinceJulianPeriod + 0.5
	new_Q := int64(Q)
	Z, _ := math.Modf(Q)
	new_Z := int64(Z)
	W := (Z - 1867216.25)/36524.25
	X := W/4
	A := Z + 1 + W - X
  	B := A + 1524
	C := (B - 122.1)/365.25
	D := 365.25 * C
	E := (B - D)/30.6001
	F := 30.6001 * E*/

	//
	fmt.Print("")

	//
	//correspondingGregorianDate.day = B - D - F

	//
	return correspondingGregorianDate
}

//
func (current_julian_day *JulianDay) Get_count_of_days_since_julian_period() float64 {

	return current_julian_day.countOfDaysSinceJulianPeriod
}

//
func (current_julian_day *JulianDay) Add_value_to_count_of_days_since_julian_period(adding_value float64) {

	current_julian_day.countOfDaysSinceJulianPeriod += adding_value
}

//
func (current_julian_day *JulianDay) Format() string {

	return fmt.Sprintf("%f", current_julian_day.countOfDaysSinceJulianPeriod)
}