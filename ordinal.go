package calendar_conversions

//
import (
	"fmt"
	"time"
)

//
type Ordinal_date struct {

	//
	year            int
	day_of_the_year int
}

//
func (current_ordinal_date *Ordinal_date) Initialize_ISO_week_based_calendar_date(dt time.Time) {

	//
	current_ordinal_date.year = dt.Year()
	current_ordinal_date.day_of_the_year = dt.YearDay()
}

//
func (current_ordinal_date *Ordinal_date) Get_year() int {

	return current_ordinal_date.year
}

//
func (current_ordinal_date *Ordinal_date) Get_day_of_the_year() int {

	return current_ordinal_date.day_of_the_year
}

//
func (current_ordinal_date *Ordinal_date) ToString() string {

	return fmt.Sprintf("%d-%d", current_ordinal_date.year, current_ordinal_date.day_of_the_year)
}
