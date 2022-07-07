package calendarConversionsGo

//
import (
    "time"
)

//
type Julian struct {

	year int64
	month time.Month
	day int64
}

//
func (current_julian *Julian) Initialize_julian(dt time.Time) {

}

//
func (current_julian *Julian) Gregorian_to_julian(gregorianDate Gregorian) {

}

//
func (current_julian *Julian) Format() string {

	return ""
}