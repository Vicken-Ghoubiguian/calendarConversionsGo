package calendarConversionsGo

import (
	"fmt"
	"time"
)

//
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
var week_day_names = [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

//
type Positivist_date struct {
}

//
func (current_positivist_date *Positivist_date) Initialize_positivist_date(dt time.Time) {
}

func main() {

	fmt.Println("hello world")
}
