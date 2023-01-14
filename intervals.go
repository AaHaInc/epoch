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
	switch i.Unit {
	case UnitSecond:
		return time.Duration(i.Value) * time.Second
	case UnitMinute:
		return time.Duration(i.Value) * time.Minute
	case UnitHour:
		return time.Duration(i.Value) * time.Hour
	case UnitDay:
		return time.Duration(i.Value) * 24 * time.Hour
	case UnitWeek:
		return time.Duration(i.Value) * 7 * 24 * time.Hour
	default:
		panic(fmt.Sprintf("unexpected unit %v", i.Unit))
	}
}

// IsSafeDuration returns true if the interval can be converted to a precise time.Duration
// This method should be used to determine if the `Duration()` method can be safely called
// on this Interval.
//
// Only seconds, minutes, hours, days, and weeks are precise.
// Interval based on months and years may be too vague and therefore
// converting them to a precise time.Duration is not possible.
func (i Interval) IsSafeDuration() bool {
	switch i.Unit {
	case UnitSecond, UnitMinute, UnitHour, UnitDay, UnitWeek:
		return true
	default:
		return false
	}
}
