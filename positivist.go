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
	Moses: {
		"a",
		"b",
		"c",
	},
	Homer: {
		"a",
		"b",
		"c",
	},
	Aristotle: {
		"a",
		"b",
		"c",
	},
	Archimedes: {
		"a",
		"b",
		"c",
	},
	Caesar: {
		"a",
		"b",
		"c",
	},
	Saint_Paul: {
		"a",
		"b",
		"c",
	},
	Charlemagne: {
		"a",
		"b",
		"c",
	},
	Dante: {
		"a",
		"b",
		"c",
	},
	Gutenberg: {
		"a",
		"b",
		"c",
	},
	Shakespeare: {
		"a",
		"b",
		"c",
	},
	Descartes: {
		"a",
		"b",
		"c",
	},
	Frederic: {
		"a",
		"b",
		"c",
	},
	Bichat: {
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
