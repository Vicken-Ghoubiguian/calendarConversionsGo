package calendarConversionsGo

//
import (
    "time"
    "fmt"
)

//
type ISO_week_based_calendar_date struct {
  
  //
  iso_year int
  iso_week_number int
  iso_week_day time.Weekday
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) InitializeWeather(dt time.Time) {
	
    //
    iso_date_s_year, iso_date_s_week := dt.ISOWeek()

    //
    current_iso_week_based_calendar_date.iso_year = iso_date_s_year
    current_iso_week_based_calendar_date.iso_week_number = iso_date_s_week
    current_iso_week_based_calendar_date.iso_week_day = dt.Weekday()
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) Get_ISO_year() int {
	
	return current_iso_week_based_calendar_date.iso_year
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) Get_ISO_week_number() int {
	
	return current_iso_week_based_calendar_date.iso_week_number
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) Get_ISO_week_day() time.Weekday {
	
	return current_iso_week_based_calendar_date.iso_week_day
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) ToString() string {

	return fmt.Sprintf("%d-W%d-%d", current_iso_week_based_calendar_date.iso_year, current_iso_week_based_calendar_date.iso_week_number, int(current_iso_week_based_calendar_date.iso_week_day))
}
