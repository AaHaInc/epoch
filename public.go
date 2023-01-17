package epoch

import (
	"time"
)

// todo: it's a draft: make it thread-safe

var globalTimeParserOptions = make([]TimeParserOption, 0)

// DefaultTimeParser is the default TimeParser that is used by public ParseTime()
var DefaultTimeParser = NewTimeParser(globalTimeParserOptions...)

func ParseTime(s string, locArg ...*time.Location) (time.Time, error) {
	return DefaultTimeParser.Parse(s, locArg...)
}

// SetIntervalArithmetics enables intervalArithmetics mode on the DefaultTimeParser
func SetIntervalArithmetics() {
	globalTimeParserOptions = append(globalTimeParserOptions, WithIntervalArithmetics())
	DefaultTimeParser = NewTimeParser(globalTimeParserOptions...)
}

func SetBaseTimeFormat(v string) {
	globalTimeParserOptions = append(globalTimeParserOptions, WithBaseTimeFormat(v))
	DefaultTimeParser = NewTimeParser(globalTimeParserOptions...)
}

func SetParsers(parsers ...Parser) {
	globalTimeParserOptions = append(globalTimeParserOptions, WithParsers(parsers...))
	DefaultTimeParser = NewTimeParser(globalTimeParserOptions...)
}
