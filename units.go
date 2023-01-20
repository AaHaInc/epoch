package epoch

type Unit struct {
	Short string
	Full  string
}

type Units []Unit

var UnitSecond = Unit{"s", "second"}
var UnitMinute = Unit{"m", "minute"}
var UnitHour = Unit{"h", "hour"}
var UnitDay = Unit{"d", "day"}
var UnitWeek = Unit{"w", "week"}
var UnitMonth = Unit{"mo", "month"}
var UnitQuarter = Unit{"q", "quarter"}
var UnitYear = Unit{"y", "year"}

var AvailableUnits = Units{UnitSecond, UnitMinute, UnitHour, UnitDay, UnitWeek, UnitMonth, UnitQuarter, UnitYear}

func (unit Unit) IsNil() bool {
	return unit.Short == ""
}

func (units Units) Get(s string) Unit {
	for _, u := range units {
		if u.Short == s {
			return u
		}
	}
	return Unit{}
}

func (units Units) Factory() *UnitFactory {
	return &UnitFactory{units}
}

// UnitFactory
// todo: refactor it, so we don't duplicate having hardcoded short units like `s`, `m`, etc
//
//	use reflectish for such kind of factories
type UnitFactory struct {
	units Units
}

func (f *UnitFactory) Second() Unit  { return f.units.Get("s") }
func (f *UnitFactory) Minute() Unit  { return f.units.Get("m") }
func (f *UnitFactory) Hour() Unit    { return f.units.Get("h") }
func (f *UnitFactory) Day() Unit     { return f.units.Get("d") }
func (f *UnitFactory) Week() Unit    { return f.units.Get("w") }
func (f *UnitFactory) Month() Unit   { return f.units.Get("mo") }
func (f *UnitFactory) Quarter() Unit { return f.units.Get("q") }
func (f *UnitFactory) Year() Unit    { return f.units.Get("y") }
