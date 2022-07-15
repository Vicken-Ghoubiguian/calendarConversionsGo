package calendarConversionsGo

//
import (
    "time"
)

//
type Gregorian struct {

	//
	year int64
	month time.Month
	monthNumber int64
	day int64
}

//


//
func (current_gregorian *Gregorian) Initialize_gregorian_from_time(dt time.Time) {

	current_gregorian.year = int64(dt.Year())
	current_gregorian.month = dt.Month()
	current_gregorian.monthNumber = int64(dt.Month())
	current_gregorian.day = int64(dt.Day())
}

//
func(current_gregorian *Gregorian) Initialize_time(year int64, monthNumber int64, day int64) time.Time {

	return time.Date(int(year), time.Month(monthNumber), int(day), 0, 0, 0, 0, time.Now().Location())
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