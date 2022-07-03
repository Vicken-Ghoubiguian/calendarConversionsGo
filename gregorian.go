package calendarConversionsGo

//
import (
    "time"
)

type gregorianMonth int

 //
 const (
	January gregorianMonth = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
 )

 //
type Gregorian struct {

	year int64
	month gregorianMonth
	day int64
}

//
func (current_gregorian *Gregorian) Initialize_gregorian(dt time.Time) {

}

 //
func (current_gregorian *Gregorian) Format() string {

	return ""
}