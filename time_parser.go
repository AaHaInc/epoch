package epoch

import (
	"fmt"
	"strings"
	"time"
)

type TimeParser struct {
	parsers                 []Parser
	withIntervalArithmetics bool
}

type TimeParserOption func(*TimeParser)

// WithIntervalArithmetics enables intervalArithmetics mode
func WithIntervalArithmetics() TimeParserOption {
	return func(tp *TimeParser) {
		tp.withIntervalArithmetics = true
	}
}

// WithBaseTimeFormat sets time format that is used by BaseParser
func WithBaseTimeFormat(v string) TimeParserOption {
	return func(tp *TimeParser) {
		BaseParserFormat = v
	}
}

// WithDefaultParsers sets the default list of parsers for TimeParser
func WithDefaultParsers() TimeParserOption {
	return func(tp *TimeParser) {
		tp.parsers = GetDefaultParsers()
	}
}

// WithParsers sets custom parsers for TimeParser
func WithParsers(parsers ...Parser) TimeParserOption {
	return func(tp *TimeParser) {
		tp.parsers = parsers
	}
}

// NewTimeParser creates a new instance of TimeParser with the provided options.
// If given options attach no parsers, it will use default parsers
func NewTimeParser(options ...TimeParserOption) *TimeParser {
	tp := &TimeParser{}
	for _, opt := range options {
		opt(tp)
	}

	if len(tp.parsers) == 0 {
		WithDefaultParsers()(tp)
	}

	return tp
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
	if !tp.withIntervalArithmetics {
		return time.Time{}, nil, fmt.Errorf("unsupported time format")
	}

	// Parse and apply each interval in given input
	for i := 1; i < len(inputs); i++ {
		interval, err := ParseInterval(inputs[i])
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("failed to parse interval [%s]: %w", inputs[i], err)
		}

		t = TimeAddInterval(t, interval)
	}

	return t, details, nil
}

// parseTime parses the given string using the list of parsers only (no interval arithmetic is applied)
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
