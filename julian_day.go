package calendarConversionsGo

//
import (
	"time"
)

//
type JulianDay struct {

	countOfDaysSinceJulianPeriod int64
}

//
func (current_julian_day *JulianDay) Initialize_julian_day(dt time.Time) {
}

//
func (current_julian_day *JulianDay) Get_count_of_days_since_julian_period(dt time.Time) int64 {

	return current_julian_day.countOfDaysSinceJulianPeriod
}

//
func (current_julian_day *JulianDay) Format() string {

	return ""
}