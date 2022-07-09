package calendarConversionsGo

type frenchRepublicanMonth int

 //
 const (
 	Vendémiaire frenchRepublicanMonth = iota
 	Brumaire
 	Frimaire
 	Nivôse
 	Pluviôse
 	Ventôse
 	Germinal
 	Floréal
 	Prairial
 	Messidor
 	Thermidor
 	Fructidor
 	Sans_culottides
 )

 type frenchRepublicanWeekDay int

 //
 const (
	Primidi frenchRepublicanWeekDay = iota
	Duodi
	Tridi
	Quartidi
	Quintidi
	Sextidi
	Septidi
	Octidi
	Nonidi
	Decadi
 )

 //
type FrenchRepublican struct {

	year int64
	month frenchRepublicanMonth
	monthDay int64
	decadi int64
	weekDay frenchRepublicanWeekDay
}