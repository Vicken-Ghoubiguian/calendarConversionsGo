package calendarConversionsGo

//
import (
	"time"
	//"strconv"
	"math"
)

//
type JulianDay struct {

	countOfDaysSinceJulianPeriod float64
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_time(dt time.Time) {

	Y := float64(dt.Year())
	M := 0.0
	D := float64(dt.Day())

	A := math.Floor(Y/100)
	B := math.Floor(A/4)
	C := math.Floor(2 - A + B)
	E := math.Floor(365.25 * (Y + 4716))
	F := math.Floor(30.6001 * (M + 1))

	JD := C + D + E + F - 1524.5

	current_julian_day.countOfDaysSinceJulianPeriod = JD
}

//
func (current_julian_day *JulianDay) Initialize_julian_day_from_gregorian_date(gregorian_date *Gregorian) {


}

//
/*func (current_julian_day *JulianDay) Determine_gregorian_date() Gregorian {


}*/

//
func (current_julian_day *JulianDay) Get_count_of_days_since_julian_period(dt time.Time) float64 {

	return current_julian_day.countOfDaysSinceJulianPeriod
}

//
func (current_julian_day *JulianDay) Format() string {

	return ""
	//return strconv.Itoa(int(current_julian_day.countOfDaysSinceJulianPeriod))
}