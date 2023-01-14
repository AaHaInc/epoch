package epoch

import (
	"fmt"
	"time"
)

type Alias struct {
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Callback    func(time.Time) time.Time
}

func GetAliasDictionary() []Alias {
	return []Alias{
		{
			Slug:        "today",
			Description: "Time of the start of today",
			Callback: func(now time.Time) time.Time {
				return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			},
		},
		{
			Slug:        "yesterday",
			Description: "Time of the start of yesterday",
			Callback: func(now time.Time) time.Time {
				return time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
			},
		},
		{
			Slug:        "tomorrow",
			Description: "Time of the start of tomorrow",
			Callback: func(now time.Time) time.Time {
				return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			},
		},
		{
			Slug:        "this-week",
			Description: "Time of the start of this week",
			Callback: func(now time.Time) time.Time {
				_, w := now.ISOWeek()
				return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -w+1)
			},
		},
		{
			Slug:        "last-week",
			Description: "Time of the start of last week",
			Callback: func(now time.Time) time.Time {
				_, w := now.AddDate(0, 0, -7).ISOWeek()
				return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -w+2)
			},
		},
		{
			Slug:        "next-week",
			Description: "Time of the start of next week",
			Callback: func(now time.Time) time.Time {
				_, w := now.AddDate(0, 0, 7).ISOWeek()
				return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -w+8)
			},
		},
	}
}

// AliasesParser parses alias strings like today, yesterday, etc
type AliasesParser struct {
	dictionary []Alias
	clock      Clock
}

var _ Parser = &AliasesParser{}

var (
	ParserNameAliases = "aliases"
)

func (a *AliasesParser) Match(s string) bool {
	for _, alias := range a.dictionary {
		if s == alias.Slug {
			return true
		}
	}
	return false
}

func (a *AliasesParser) Parse(s string, locArg ...*time.Location) (time.Time, *ParseDetails, error) {
	var loc *time.Location
	if len(locArg) > 0 {
		loc = locArg[0]
	}

	for _, alias := range a.dictionary {
		if s != alias.Slug {
			continue
		}

		now := a.clock.Now()
		if loc != nil {
			now = now.In(loc)
		}
		return alias.Callback(now), &ParseDetails{
			IsRelative: true,
			IsAliased:  true,
			ParserName: ParserNameAliases,
		}, nil
	}

	return time.Time{}, nil, fmt.Errorf("alias not found")
}

func (a *AliasesParser) GetDictionary() []Alias {
	return a.dictionary
}

func (a *AliasesParser) ExpandDictionary(aliases ...Alias) {
	a.dictionary = append(a.dictionary, aliases...)
}

func (a *AliasesParser) SetClock(c Clock) *AliasesParser {
	a.clock = c
	return a
}

func NewAliasesParser() *AliasesParser {
	return &AliasesParser{
		dictionary: GetAliasDictionary(),
		clock:      NewDefaultClock(),
	}
}

// Name returns the name of the parser, "aliases"
func (u *AliasesParser) Name() string {
	return ParserNameAliases
}
