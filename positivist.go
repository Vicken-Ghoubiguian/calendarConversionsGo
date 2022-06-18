package calendarConversionsGo

import (
	"fmt"
	"time"
)

//
var days_from_months = map[string][]string{
	"1": {
		"a",
		"b",
		"c",
	},
	"2": {
		"a",
		"b",
		"c",
	},
	"3": {
		"a",
		"b",
		"c",
	},
}

//
type Positivist_date struct {
}

//
func (current_positivist_date *Positivist_date) Initialize_positivist_date(dt time.Time) {
}

func main() {

	fmt.Println("hello world")
}
