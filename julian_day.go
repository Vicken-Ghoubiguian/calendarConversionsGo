package calendarConversionsGo

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
	current_julian_day.countOfDaysSinceJulianPeriod = C + D + E + F - 1524.5
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_gregorian_date(gregorian_date *Gregorian) {

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
	current_julian_day.countOfDaysSinceJulianPeriod = C + D + E + F - 1524.5
}

//
func (current_julian_day *JulianDay) Determine_gregorian_date() {


}

//
func (current_julian_day *JulianDay) Get_count_of_days_since_julian_period(dt time.Time) float64 {

	return current_julian_day.countOfDaysSinceJulianPeriod
}

//
func (current_julian_day *JulianDay) Format() string {

	return fmt.Sprintf("%f", current_julian_day.countOfDaysSinceJulianPeriod)
}