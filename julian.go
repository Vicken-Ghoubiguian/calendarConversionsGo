package calendar_conversions

//
import (
    "time"
	"fmt"
	"math"
)

//
type Julian struct {

	year int64
	month time.Month
	monthNumber int64
	day int64
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
	current_julian.day = int64(day)
	current_julian.month = time.Month(month)
	current_julian.monthNumber = int64(month)
	current_julian.year = int64(year)
}

//
func (current_julian *Julian) Julian_to_Gregorian() Gregorian {

	//
	var gd Gregorian

	return gd
}

//
func (current_julian *Julian) Get_year() int64 {

	return current_julian.year
}

//
func (current_julian *Julian) Get_month() time.Month {

	return current_julian.month
}

//
func (current_julian *Julian) Get_monthNumber() int64 {

	return current_julian.monthNumber
}

//
func (current_julian *Julian) Get_day() int64 {

	return current_julian.day
}

//
func (current_julian *Julian) Format() string {

	return fmt.Sprintf("%d-%d-%d", current_julian.year, int(current_julian.month), current_julian.day)
}