package epoch

import (
	"errors"
	"fmt"
	"strings"
)

type Interval struct {
	Second, Minute, Hour, Day, Week, Month, Year float64
}

func ParseInterval(interval string) (*Interval, error) {
	interval = strings.TrimSpace(interval)
	if interval == "" {
		return nil, errors.New("interval string cannot be empty")
	}

	var i Interval
	var value float64
	var unit string

	n, err := fmt.Sscanf(interval, "%f%s", &value, &unit)
	if err != nil || n != 2 {
		return nil, fmt.Errorf("invalid interval string format: %s", interval)
	}

	switch strings.ToLower(unit) {
	case "s", "sec", "second":
		i.Second = value
	case "m", "min", "minute":
		i.Minute = value / 60
	case "h", "hour":
		i.Hour = value / 3600
	case "d", "day":
		i.Day = value / 86400
	case "w", "week":
		i.Week = value / 604800
	case "month":
		i.Month = value / 2592000
	case "y", "year":
		i.Year = value / 31536000
	default:
		return nil, fmt.Errorf("invalid interval unit: %s", unit)
	}

	return &i, nil
}

func main() {
	interval, err := ParseInterval("5m")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Interval: %+v\n", interval)
}
