package main

import (
	"errors"
	"fmt"
	"strings"
)

type Interval struct {
	Value float64
	Unit  string
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
	case "s":
		i.Unit = "second"
	case "m":
		i.Unit = "minute"
	case "h":
		i.Unit = "hour"
	case "d":
		i.Unit = "day"
	case "w":
		i.Unit = "week"
	case "mo":
		i.Unit = "month"
	case "y":
		i.Unit = "year"
	default:
		return nil, fmt.Errorf("invalid interval unit: %s", unit)
	}
	i.Value = value
	return &i, nil
}

func main() {
	interval, err := ParseInterval("5mo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Interval: %+v\n", interval)
}
