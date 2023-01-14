package epoch_test

import (
	"github.com/aahainc/epoch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("TimeParser", func() {
	Context("With default parsers", func() {
		p := epoch.NewTimeParser()

		When("valid inputs were given (no timezone is given)", func() {
			DescribeTable("absolute dates", func(input string, tExpected time.Time) {
				t, err := p.Parse(input)
				Expect(err).Should(Succeed())
				Expect(t.UnixMilli()).To(Equal(tExpected.UnixMilli()))
			},
				Entry("unix timestamp in seconds", "1577829600", time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)),
				Entry("a simple date (RFC3339 format)", "2018-03-01T12:30:00Z", time.Date(2018, time.March, 1, 12, 30, 0, 0, time.UTC)),
				Entry("RFC3339 format with fractional seconds", "2018-03-01T12:30:00.123Z", time.Date(2018, time.March, 1, 12, 30, 0, 123000000, time.UTC)),
				Entry("RFC3339 format with timezone offset", "2018-03-01T12:30:00-07:00", time.Date(2018, time.March, 1, 12, 30, 0, 0, time.FixedZone("-07:00", -25200))),
				Entry("RFC3339 format with timezone offset and fractional seconds", "2018-03-01T12:30:00.456-07:00", time.Date(2018, time.March, 1, 12, 30, 0, 456000000, time.FixedZone("-07:00", -25200))),
			)
		})

		When("mocked time.Now() to a fixed date", func() {
			var p *epoch.TimeParser

			fixedNow := time.Date(2006, time.January, 02, 15, 4, 5, 0, time.UTC)
			parsersWithFixedNow := []epoch.Parser{
				epoch.NewRFC3339Parser(),
				epoch.NewUnixSecondsParser(),
				epoch.NewAliasesParser().SetClock(epoch.NewStaticClock(fixedNow)),
			}

			BeforeEach(func() {
				p = epoch.NewTimeParser(parsersWithFixedNow...)
			})

			DescribeTable("relative dates", func(input string, tExpected time.Time) {
				t, err := p.Parse(input, time.UTC)
				Expect(err).Should(Succeed())
				Expect(t.UnixMilli()).To(Equal(tExpected.UnixMilli()))
			},
				Entry("today", "today", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
				Entry("yesterday", "yesterday", time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Entry("tomorrow", "tomorrow", time.Date(2006, time.January, 3, 0, 0, 0, 0, time.UTC)),
				//Entry("this week", "this-week", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
				//Entry("last week", "last-week", time.Date(2005, time.December, 26, 0, 0, 0, 0, time.UTC)),
				//Entry("next week", "next-week", time.Date(2006, time.January, 9, 0, 0, 0, 0, time.UTC)),
				//Entry("this month", "this-month", time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)),
				//Entry("last month", "last-month", time.Date(2005, time.December, 1, 0, 0, 0, 0, time.UTC)),
				//Entry("next month", "next-month", time.Date(2006, time.February, 1, 0, 0, 0, 0, time.UTC)),
				//Entry("this year", "this-year", time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)),
				//Entry("last year", "last-year", time.Date(2005, time.January, 1, 0, 0, 0, 0, time.UTC)),
				//Entry("next year", "next-year", time.Date(2007, time.January, 1, 0, 0, 0, 0, time.UTC)),
			)

			When("and with arithmetics", func() {
				BeforeEach(func() {
					p.IntervalArithmeticsEnabled = true
				})

				DescribeTable("relative dates", func(input string, tExpected time.Time) {
					t, err := p.Parse(input, time.UTC)
					Expect(err).Should(Succeed())
					Expect(t.UnixMilli()).To(Equal(tExpected.UnixMilli()))
				},
					Entry("the day before today", "today,-1d", time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)),
				)
			})
		})

	})
})
