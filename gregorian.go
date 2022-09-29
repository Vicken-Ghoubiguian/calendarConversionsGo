package calendar_conversions

//
import (
    "time"
	"fmt"
	"strconv"
)

//
type Gregorian struct {

	//
	year int64
	month time.Month
	monthNumber int64
	day int64
	hour int64
	minute int64
	second int64
	microseconds int64
}

//
func (current_gregorian *Gregorian) Initialize_gregorian_from_elements(year int64, monthNumber int64, day int64, hour int64, minute int64, second int64, microseconds int64) {

	current_gregorian.year = year
	current_gregorian.month = time.Month(monthNumber)
	current_gregorian.monthNumber = monthNumber
	current_gregorian.day = day
	current_gregorian.hour = hour
	current_gregorian.minute = minute
	current_gregorian.second = second
	//current_gregorian.microseconds = microseconds
}

//
func (current_gregorian *Gregorian) Initialize_gregorian_from_time(dt time.Time) {

	current_gregorian.year = int64(dt.Year())
	current_gregorian.month = dt.Month()
	current_gregorian.monthNumber = int64(dt.Month())
	current_gregorian.day = int64(dt.Day())
	current_gregorian.hour = int64(dt.Hour())
	current_gregorian.minute = int64(dt.Minute())
	current_gregorian.second = int64(dt.Second())
	//current_gregorian.microseconds = int64(dt.Microsecond())
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
func (current_gregorian *Gregorian) Get_hour() int64 {

	return current_gregorian.hour
}

//
func (current_gregorian *Gregorian) Get_minute() int64 {

	return current_gregorian.minute
}

//
func (current_gregorian *Gregorian) Get_second() int64 {

	return current_gregorian.second
}

//
/*func (current_gregorian *Gregorian) Get_microseconds() int64 {

	return current_gregorian.microseconds
}*/

//
func (current_gregorian *Gregorian) Format() string {

	year := strconv.FormatInt(current_gregorian.year, 10)

	month := strconv.FormatInt(int64(current_gregorian.month), 10)
	if int64(current_gregorian.month) < 10 {

		month = "0" + month
	}

	day := "0" + strconv.FormatInt(current_gregorian.day, 10)
	if current_gregorian.day < 10 {

		day = "0" + day
	}

	return fmt.Sprintf("%s-%s-%s", year, month, day)
}