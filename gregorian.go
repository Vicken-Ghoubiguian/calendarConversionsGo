package calendarConversionsGo

//
import (
    "time"
)

//
type Gregorian struct {

	year int64
	month gregorianAndJulianMonth
	monthNumber int64
	day int64
}

//
func (current_gregorian *Gregorian) Initialize_gregorian(dt time.Time) {

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
func (current_gregorian *Gregorian) Get_month() gregorianAndJulianMonth {

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