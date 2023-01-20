package epoch

import (
	"fmt"
	"strconv"
	"time"
)

var (
	ErrInvalidUnit   = fmt.Errorf("invalid unit")
	ErrInvalidFormat = fmt.Errorf("invalid format")
)

type Interval struct {
	Value float64
	Unit  Unit
}

func ParseInterval(interval string) (*Interval, error) {
	var value float64
	var unitShort string
	n, err := fmt.Sscanf(interval, "%f%s", &value, &unitShort)
	if n != 2 || err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidFormat, err)
	}

	unit := AvailableUnits.Get(unitShort)
	if unit.IsNil() {
		return nil, ErrInvalidUnit
	}

	return &Interval{Value: value, Unit: unit}, nil
}

func MustParseInterval(interval string) *Interval {
	i, err := ParseInterval(interval)
	if err != nil {
		panic(err)
	}
	return i
}

func (i *Interval) String() string {
	return strconv.FormatFloat(i.Value, 'f', -1, 64) + i.Unit.Short
}

// IsNil returns true if interval is nil
func (i *Interval) IsNil() bool {
	if i == nil {
		return true
	}

	return i.Unit.IsNil()
}

// IsSafeDuration returns true if the interval can be converted to a precise time.Duration
// This method should be used to determine if the `Duration()` method can be safely called
// on this Interval.
//
// Only seconds, minutes, hours, days, and weeks are precise.
// Interval based on months and years may be too vague and therefore
// converting them to a precise time.Duration is not possible.
func (i *Interval) IsSafeDuration() bool {
	switch i.Unit {
	case UnitSecond, UnitMinute, UnitHour, UnitDay, UnitWeek:
		return true
	default:
		return false
	}
}

// Duration returns time.Duration when it's safe (see IsSafeDuration)
// It will panic otherwise
func (i *Interval) Duration() time.Duration {
	switch i.Unit {
	case UnitSecond:
		return time.Duration(i.Value * float64(time.Second))
	case UnitMinute:
		return time.Duration(i.Value * float64(time.Minute))
	case UnitHour:
		return time.Duration(i.Value * float64(time.Hour))
	case UnitDay:
		return time.Duration(i.Value * float64(24*time.Hour))
	case UnitWeek:
		return time.Duration(i.Value * float64(7*24*time.Hour))
	default:
		panic(fmt.Sprintf("can't get duration of %v", i.Unit))
	}
}

// ExtractDateParts returns the year, month, and day as integers of an Interval.
// It's considered to be used to add the interval to a time.Time using time.AddDate()
// ExtractDateParts returns the number of years, months, and days in the interval.
// It can be used in conjunction with time.Time.AddDate to move a time.Time by the duration of the interval.
func (i *Interval) ExtractDateParts() (years int, months int, days int) {
	switch i.Unit {
	case UnitYear:
		years = int(i.Value)
	case UnitMonth:
		months = int(i.Value)
	case UnitWeek:
		days = int(i.Value * 7)
	case UnitDay:
		days = int(i.Value)
	}
	return
}
