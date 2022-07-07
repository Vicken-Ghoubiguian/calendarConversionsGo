package calendarConversionsGo

//
import (
    "time"
)

//
type Gregorian struct {

	year int64
	month time.Month
	monthNumber int64
	day int64
}

//
func (current_gregorian *Gregorian) Initialize_gregorian(dt time.Time) {

	current_gregorian.year = int64(dt.Year())
	current_gregorian.month = dt.Month()
	current_gregorian.day = int64(dt.Day())
}

//
func (current_gregorian *Gregorian) Julian_to_gregorian(julianDate Julian) {

}

//
func (current_gregorian *Gregorian) JulianDay_to_gregorian(julianDayDate JulianDay) {

}

//
func (current_gregorian *Gregorian) Get_year() int64 {

	return current_gregorian.year
}

//
func (current_gregorian *Gregorian) Get_month() time.Month {

	return current_gregorian.month
}

//
func (current_gregorian *Gregorian) Get_monthNumber() int64 {

	return current_gregorian.monthNumber
}

//
func (current_gregorian *Gregorian) Get_day() int64 {

	return current_gregorian.day
}

//
func (current_gregorian *Gregorian) Format() string {

	return ""
}