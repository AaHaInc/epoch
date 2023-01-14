package epoch

import (
	"fmt"
	"time"
)

// RFC3339Parser parses time in RFC3339 format
type RFC3339Parser struct{}

var _ Parser = &RFC3339Parser{}

// NewRFC3339Parser returns a new RFC3339Parser
func NewRFC3339Parser() *RFC3339Parser {
	return &RFC3339Parser{}
}

// Match checks if given string is in RFC3339 format
func (r *RFC3339Parser) Match(s string) bool {
	_, err := time.Parse(time.RFC3339, s)
	return err == nil
}

// Parse converts string to time.Time
func (r *RFC3339Parser) Parse(s string, locArg ...*time.Location) (time.Time, error) {
	loc := time.UTC
	if len(locArg) > 0 {
		loc = locArg[0]
	}

	t, err := time.ParseInLocation(time.RFC3339, s, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse RFC3339 time: %w", err)
	}
	return t, nil
}
