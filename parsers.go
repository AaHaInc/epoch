package epoch

import (
	"time"
)

// Parser is an interface for parsing time string
type Parser interface {
	// Match checks if parser supports given string
	Match(s string) bool
	// Parse converts string to time.Time and returns a ParseDetails struct
	// that includes details of the parsing process
	Parse(s string, locArg ...*time.Location) (time.Time, *ParseDetails, error)
	// Name returns the name of the parser
	Name() string
}

// ParseDetails stores details of parsing.
type ParseDetails struct {
	// ParserName is the name of the parser that was chosen
	ParserName string `json:"parser_name"`
	// IsRelative indicates whether the given time is relative or absolute
	IsRelative bool `json:"is_relative"`
	// IsAliased indicates whether the given time was taken from an alias or not
	IsAliased bool `json:"is_aliased"`
	// Format stores the format used for parsing a formatted time.
	Format string `json:"format"`
	// Arithmetics stores information about arithmetic operations applied to parsed time
	Arithmetics *Arithmetics `json:"arithmetics,omitempty"`
}

// Arithmetics holds information about any arithmetic operations applied on a parsed time
type Arithmetics struct {
	// Intervals is a list of intervals used in arithmetic operations
	Intervals []Interval `json:"intervals"`
	// RawIntervals is a list of the raw interval strings used in arithmetic operations
	RawIntervals []string `json:"raw_intervals"`
}

// GetDefaultParsers returns a list of default parsers including RFC3339, Unix Seconds and Aliases Parsers
func GetDefaultParsers() []Parser {
	return []Parser{
		NewBaseParser(),
		NewUnixSecondsParser(),
		NewAliasesParser(),
	}
}

// GetAllParsers returns a list of all the parsers we support
func GetAllParsers() []Parser {
	return []Parser{
		NewBaseParser(),
		NewUnixMilliParser(),
		NewUnixSecondsParser(),
		NewAliasesParser(),
	}
}
