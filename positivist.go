package calendarConversionsGo

import (
	"fmt"
	"time"
)

type Month int

 //
 const (
 	Moses Month = iota
 	Homer
 	Aristotle
 	Archimedes
 	Caesar
 	Saint_Paul
 	Charlemagne
 	Dante
 	Gutenberg
 	Shakespeare
 	Descartes
 	Frederic
 	Bichat
 )

//
var days_from_months = map[Month][]string{
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
	"1": {
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
