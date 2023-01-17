package epoch

import (
	"fmt"
	"time"
)

// BaseParser parses time in a specified format (defaulted to time.RFC3339)
type BaseParser struct{}

var _ Parser = &BaseParser{}

var (
	ParserNameBase = "base"
)

var (
	// BaseParserFormat is the time format used for BaseParser
	// Can be overwritten via WithBaseTimeFormat()
	BaseParserFormat = time.RFC3339
)

// NewBaseParser returns a new BaseParser with the specified format
func NewBaseParser() *BaseParser {
	return &BaseParser{}
}

// Match checks if given string is in the specified format
func (b *BaseParser) Match(s string) bool {
	_, err := time.Parse(BaseParserFormat, s)
	return err == nil
}

// Parse converts string to time.Time
func (b *BaseParser) Parse(s string, locArg ...*time.Location) (time.Time, *ParseDetails, error) {
	var loc *time.Location
	if len(locArg) > 0 {
		loc = locArg[0]
	}

	var t time.Time
	var err error
	if loc != nil {
		t, err = time.ParseInLocation(BaseParserFormat, s, loc)
	} else {
		t, err = time.Parse(BaseParserFormat, s)
	}
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("failed to parse time in format %s: %w", BaseParserFormat, err)
	}

	// if no location is given
	// load the location from the parsed time
	if loc == nil && t.Location() != nil {
		loc, err = time.LoadLocation(t.Location().String())
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("invalid location specified: %w", err)
		}

		t, err = time.ParseInLocation(BaseParserFormat, s, loc)
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("failed to parse time in location in format %s: %w", BaseParserFormat, err)
		}
	}

	return t, &ParseDetails{
		ParserName: ParserNameBase,
		Format:     BaseParserFormat,
	}, nil
}

// Name returns the name of the parser, "base"
func (b *BaseParser) Name() string {
	return ParserNameBase
}
