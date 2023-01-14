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
var UnitYear = Unit{"y", "year"}

var AvailableUnits = Units{UnitSecond, UnitMinute, UnitHour, UnitDay, UnitWeek, UnitMonth, UnitYear}

func (units Units) Get(s string) *Unit {
	for _, u := range units {
		if u.Short == s {
			return &u
		}
	}
	return nil
}

func ListAvailableUnits() Units {
	return AvailableUnits
}
