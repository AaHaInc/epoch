package epoch

import (
	"time"
)

// Parser is an interface for parsing time string
type Parser interface {
	// Match checks if parser supports given string
	Match(s string) bool
	// Parse converts string to time.Time
	Parse(s string, locArg ...*time.Location) (time.Time, error)
}

// GetDefaultParsers returns a list of default parsers including RFC3339, Unix Seconds and Aliases Parsers
func GetDefaultParsers() []Parser {
	return []Parser{
		NewRFC3339Parser(),
		NewUnixSecondsParser(),
		NewAliasesParser(),
	}
}

// GetAllParsers returns a list of all the parsers we support
func GetAllParsers() []Parser {
	return []Parser{
		NewRFC3339Parser(),
		NewUnixMilliParser(),
		NewUnixSecondsParser(),
		NewAliasesParser(),
	}
}
