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
type iso_date struct {
  
  //
  iso_year int
  iso_week_number int
  iso_week_day int
}

//
func (current_iso_date *iso_date) InitializeWeather(dt time) {

	//
    current_iso_date.iso_year = 0
    current_iso_date.iso_week_number = 0
    current_iso_date.iso_week_day = 0
}
