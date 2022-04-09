package calendarConversionsGo

//
import (
    "time"
)

//
const (
        sunday = 0
        monday = 1
        tuesday = 2
        wednesday = 3
        thursday = 4
        friday = 5
        saturday = 6
)

//
type ISO_date struct {
  
  //
  iso_year int
  iso_week_number int
  iso_week_day int
}

//
func (current_iso_date *ISO_date) InitializeWeather(dt time) {
	
    //
    iso_date_s_year, iso_date_s_week := t.ISOWeek()

    //
    current_iso_date.iso_year = iso_date_s_year
    current_iso_date.iso_week_number = iso_date_s_week
    current_iso_date.iso_week_day = t.Weekday()
}
