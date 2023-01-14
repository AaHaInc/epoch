package epoch

import (
	"fmt"
	"strconv"
	"time"
)

// UnixMilliParser parses unix timestamp in milliseconds
type UnixMilliParser struct{}

var _ Parser = &UnixMilliParser{}

// NewUnixMilliParser returns a new UnixMilliParser
func NewUnixMilliParser() *UnixMilliParser {
	return &UnixMilliParser{}
}

// Match checks if given string is unix timestamp in milliseconds
func (u *UnixMilliParser) Match(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil && len(s) >= 11
}

// Parse converts string to time.Time
func (u *UnixMilliParser) Parse(s string, locArg ...*time.Location) (time.Time, error) {
	loc := time.UTC
	if len(locArg) > 0 {
		loc = locArg[0]
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse unix milliseconds time: %w", err)
	}
	return time.Unix(0, i*int64(time.Millisecond)).In(loc), nil
}
