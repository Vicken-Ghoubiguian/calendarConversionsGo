package calendarConversionsGo

//
import (
    "time"
    "fmt"
)

//
type ISO_date struct {
  
  //
  iso_year int
  iso_week_number int
  iso_week_day time.Weekday
}

//
func (current_iso_date *ISO_date) InitializeWeather(dt time.Time) {
	
    //
    iso_date_s_year, iso_date_s_week := dt.ISOWeek()

    //
    current_iso_date.iso_year = iso_date_s_year
    current_iso_date.iso_week_number = iso_date_s_week
    current_iso_date.iso_week_day = dt.Weekday()
}

//
func (current_iso_date *ISO_date) Get_ISO_year() int {
	
	return current_iso_date.iso_year
}

//
func (current_iso_date *ISO_date) Get_ISO_week_number() int {
	
	return current_iso_date.iso_week_number
}

//
func (current_iso_date *ISO_date) ToString() string {

	return fmt.Sprintf("%d-W%d-%d", current_iso_date.iso_year, int(current_iso_date.iso_week_number), current_iso_date.iso_week_number)
}

//
func (current_iso_date *ISO_date) Get_ISO_week_day() time.Weekday {
	
	return current_iso_date.iso_week_day
}
