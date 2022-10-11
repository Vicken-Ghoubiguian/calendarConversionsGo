package calendar_conversions

//
import (
    "time"
	"fmt"
	"math"
	"strconv"
)

//
type Julian struct {

	year int
	month time.Month
	monthNumber int
	day int
	hour int
	minute int
	second int
	microseconds int
}

//
func (current_julian *Julian) Initialize_julian_from_Gregorian(gregorianDate Gregorian) {

	//
	var jd JulianDay

	//
	jd.Initialize_julian_day_from_gregorian_date(gregorianDate)

	//
	jd.Add_value_to_count_of_days_since_julian_period(0.5)
	a := math.Floor(jd.Get_count_of_days_since_julian_period())
	b := a + 1524
	c := math.Floor((b - 122.1) / 365.25)
	d := math.Floor(365.25 * c)
	e := math.Floor((b - d) / 30.6001)
	
	//
	month := -1

	//
	if e < 14 {

		month = int(math.Floor(e - 1))

	} else {

		month = int(math.Floor(e - 13))
	}

	//
	year := -1

	//
	if month > 2 {

		year = int(math.Floor(c - 4716))

	} else {

		year = int(math.Floor(c - 4715))
	}

	//
	day := b - d - math.Floor(30.6001 * e)
	
	//
	current_julian.day = int(day)
	current_julian.month = time.Month(month)
	current_julian.monthNumber = int(month)
	current_julian.year = int(year)

	//
	current_julian.hour = 0
	current_julian.minute = 0
	current_julian.second = 0
	current_julian.microseconds = 0
}

//
func (current_julian *Julian) Julian_to_Gregorian() Gregorian {

	//
	var gd Gregorian

	//
	return gd
}

//
func (current_julian *Julian) Get_year() int {

	return current_julian.year
}

//
func (current_julian *Julian) Get_month() time.Month {

	return current_julian.month
}

//
func (current_julian *Julian) Get_monthNumber() int {

	return current_julian.monthNumber
}

//
func (current_julian *Julian) Get_day() int {

	return current_julian.day
}

//
func (current_julian *Julian) Get_hour() int {

	return current_julian.hour
}

//
func (current_julian *Julian) Get_minute() int {

	return current_julian.minute
}

//
func (current_julian *Julian) Get_second() int {

	return current_julian.second
}

//
func (current_julian *Julian) Get_microseconds() int {

	return current_julian.microseconds
}

//
func (current_julian *Julian) Format() string {

	year := strconv.Itoa(current_julian.year)

	month := strconv.Itoa(int(current_julian.month))
	if int(current_julian.month) < 10 {

		month = "0" + month
	}

	day := strconv.Itoa(current_julian.day)
	if current_julian.day < 10 {

		day = "0" + day
	}

	hour := strconv.Itoa(current_julian.hour)
	if current_julian.hour < 10 {

		hour = "0" + hour
	}

	minute := strconv.Itoa(current_julian.minute)
	if current_julian.minute < 10 {

		minute = "0" + minute
	}

	second := strconv.Itoa(current_julian.second)
	if current_julian.second < 10 {

		second = "0" + second
	}

	microseconds := strconv.Itoa(current_julian.microseconds)

	return fmt.Sprintf("%s-%s-%s %s:%s:%s.%s", year, month, day, hour, minute, second, microseconds)
}