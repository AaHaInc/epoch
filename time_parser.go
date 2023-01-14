package epoch

import (
	"fmt"
	"strings"
	"time"
)

type TimeParser struct {
	parsers                    []Parser
	IntervalArithmeticsEnabled bool
}

// Parse attempts to parse the given string using the list of parsers.
func (tp *TimeParser) Parse(s string, locArg ...*time.Location) (time.Time, error) {
	t, _, err := tp.ParseExt(s, locArg...)
	return t, err
}

// ParseExt attempts to parse the given string using the list of parsers.
func (tp *TimeParser) ParseExt(s string, locArg ...*time.Location) (time.Time, *ParseDetails, error) {
	inputs := strings.Split(s, ",")
	t, details, err := tp.parseTime(inputs[0], locArg...)
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("failed to parse time: %w", err)
	}

	// if we have just one input, not
	if len(inputs) == 1 {
		return t, details, nil
	}

	// comma-separated value are allowed only if interval arithmetics is enabled
	if !tp.IntervalArithmeticsEnabled {
		return time.Time{}, nil, fmt.Errorf("unsupported time format")
	}

	for i := 1; i < len(inputs); i++ {
		interval, err := ParseInterval(inputs[i])
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("failed to parse interval [%s]: %w", inputs[i], err)
		}

		t = TimeAddInterval(t, interval)
	}

	return t, details, nil
}

func NewTimeParser(parsers ...Parser) *TimeParser {
	if len(parsers) == 0 {
		parsers = GetDefaultParsers()
	}
	return &TimeParser{parsers: parsers}
}

func (tp *TimeParser) parseTime(s string, locArg ...*time.Location) (time.Time, *ParseDetails, error) {
	for _, parser := range tp.parsers {
		if !parser.Match(s) {
			continue
		}
		t, details, err := parser.Parse(s, locArg...)
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("failed to parse time: %w", err)
		}
		return t, details, nil
	}
	return time.Time{}, nil, fmt.Errorf("unsupported time format")
}
