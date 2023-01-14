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
	inputs := strings.Split(s, ",")
	t, err := tp.parseTime(inputs[0], locArg...)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %w", err)
	}

	// if we have just one input, not
	if len(inputs) == 1 {
		return t, nil
	}

	// comma-separated value are allowed only if interval arithmetics is enabled
	if !tp.IntervalArithmeticsEnabled {
		return time.Time{}, fmt.Errorf("unsupported time format")
	}

	for i := 1; i < len(inputs); i++ {
		interval, err := ParseInterval(inputs[i])
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to parse interval [%s]: %w", inputs[i], err)
		}

		t = TimeAddInterval(t, interval)
	}

	return t, nil
}

func NewTimeParser(parsers ...Parser) *TimeParser {
	if len(parsers) == 0 {
		parsers = GetDefaultParsers()
	}
	return &TimeParser{parsers: parsers}
}

func (tp *TimeParser) parseTime(s string, locArg ...*time.Location) (time.Time, error) {
	for _, parser := range tp.parsers {
		if !parser.Match(s) {
			continue
		}
		t, _, err := parser.Parse(s, locArg...)
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to parse time: %w", err)
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("unsupported time format")
}
