package epoch

import (
	"fmt"
	"strconv"
	"time"
)

// UnixSecondsParser parses unix timestamp in seconds
type UnixSecondsParser struct{}

var _ Parser = &UnixSecondsParser{}

var (
	ParserNameUnixSeconds = "unix-seconds"
)

// NewUnixSecondsParser returns a new UnixSecondsParser
func NewUnixSecondsParser() *UnixSecondsParser {
	return &UnixSecondsParser{}
}

// Match checks if given string is unix timestamp in seconds
func (u *UnixSecondsParser) Match(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil && len(s) < 11
}

// Parse converts string to time.Time
func (u *UnixSecondsParser) Parse(s string, locArg ...*time.Location) (time.Time, *ParseDetails, error) {
	loc := time.UTC
	if len(locArg) > 0 {
		loc = locArg[0]
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("failed to parse unix seconds time: %w", err)
	}
	return time.Unix(i, 0).In(loc), &ParseDetails{
		ParserName: ParserNameUnixSeconds,
	}, nil
}

// Name returns the name of the parser, "unix-seconds"
func (u *UnixSecondsParser) Name() string {
	return ParserNameUnixSeconds
}
