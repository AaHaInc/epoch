package epoch

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Interval struct {
	Unit  Unit
	Value float64
}

func ParseInterval(interval string) (*Interval, error) {
	interval = strings.TrimSpace(interval)
	if len(interval) < 2 {
		return nil, fmt.Errorf("interval string is too short")
	}
	var value float64
	var unitStr string
	n, err := fmt.Sscanf(interval, "%f%s", &value, &unitStr)
	if n != 2 || err != nil {
		return nil, fmt.Errorf("failed to parse interval string: %s", err)
	}
	unit := ListAvailableUnits().Get(unitStr)
	if unit == nil {
		return nil, fmt.Errorf("unknown unit %q", unitStr)
	}
	return &Interval{Unit: *unit, Value: value}, nil
}

func MustParseInterval(interval string) *Interval {
	i, err := ParseInterval(interval)
	if err != nil {
		panic(err)
	}
	return i
}

func (i Interval) String() string {
	return strconv.FormatFloat(i.Value, 'f', -1, 64) + i.Unit.Short
}

func (i Interval) Duration() time.Duration {
	var d time.Duration
	switch i.Unit {
	case UnitSecond:
		d = time.Duration(i.Value * float64(time.Second))
	case UnitMinute:
		d = time.Duration(i.Value * float64(time.Minute))
	case UnitHour:
		d = time.Duration(i.Value * float64(time.Hour))
	case UnitDay:
		d = time.Duration(i.Value * 24 * float64(time.Hour))
	case UnitWeek:
		d = time.Duration(i.Value * 7 * 24 * float64(time.Hour))
	case UnitMonth:
		d = time.Duration(i.Value * 30 * 24 * float64(time.Hour))
	case UnitYear:
		d = time.Duration(i.Value * 365 * 24 * float64(time.Hour))
	}
	return d
}
