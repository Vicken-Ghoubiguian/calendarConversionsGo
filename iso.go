package calendarConversionsGo

//
import (
    "fmt"
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
  iso_year int64
  iso_week_number int64
  iso_week_day int64
}
