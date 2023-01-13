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

func getUnits() map[string]string {
	return map[string]string{
		"s":  "second",
		"m":  "minute",
		"h":  "hour",
		"d":  "day",
		"w":  "week",
		"mo": "month",
		"y":  "year",
	}
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
	units := getUnits()
	if _, ok := units[strings.ToLower(unit)]; !ok {
		return nil, fmt.Errorf("invalid interval unit: %s", unit)
	}
	i.Unit = units[strings.ToLower(unit)]
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
