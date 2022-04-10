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
    iso_week_day int
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) Initialize_ISO_week_based_calendar_date(dt time.Time) {
      
    //
    iso_date_s_year, iso_date_s_week := dt.ISOWeek()
  
    //
    current_iso_week_based_calendar_date.iso_year = iso_date_s_year
    current_iso_week_based_calendar_date.iso_week_number = iso_date_s_week

    if int(dt.Weekday()) == 0 {
        current_iso_week_based_calendar_date.iso_week_day = 7

    } else {
        current_iso_week_based_calendar_date.iso_week_day = int(dt.Weekday())
    }
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
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) Get_ISO_week_day() int {
	
	return current_iso_week_based_calendar_date.iso_week_day
}

//
func (current_iso_week_based_calendar_date *ISO_week_based_calendar_date) ToString() string {

	return fmt.Sprintf("%d-W%d-%d", current_iso_week_based_calendar_date.iso_year, current_iso_week_based_calendar_date.iso_week_number, current_iso_week_based_calendar_date.iso_week_day)
}
