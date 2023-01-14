package epoch_test

import (
	"github.com/aahainc/epoch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Helpers", func() {

	Context("Rounding", func() {
		Context("Truncate to hour (rounding down)", func() {
			It("should truncate to hour (UTC)", func() {
				a := time.Date(2019, 10, 12, 5, 32, 0, 0, time.UTC)
				result := epoch.TruncateToHour(a)
				Expect(result.String()).To(Equal("2019-10-12 05:00:00 +0000 UTC"))
			})

			It("should truncate to hour (location)", func() {
				loc, _ := time.LoadLocation("America/Los_Angeles")
				a := time.Date(2019, 10, 12, 5, 32, 0, 0, loc)
				result := epoch.TruncateToHour(a)
				Expect(result.String()).To(Equal("2019-10-12 05:00:00 -0700 PDT"))
			})

			It("should truncate to hour (edge case: last minute of hour)", func() {
				a := time.Date(2019, 10, 12, 5, 59, 0, 0, time.UTC)
				result := epoch.TruncateToHour(a)
				Expect(result.String()).To(Equal("2019-10-12 05:00:00 +0000 UTC"))
			})

			It("should truncate to hour (edge case: first minute of hour)", func() {
				a := time.Date(2019, 10, 12, 5, 0, 0, 0, time.UTC)
				result := epoch.TruncateToHour(a)
				Expect(result.String()).To(Equal("2019-10-12 05:00:00 +0000 UTC"))
			})
		})

	})

	Context("EffectiveHoursInDay", func() {
		It("should return 24 hours for most days", func() {
			loc, _ := time.LoadLocation("America/Los_Angeles")
			t := time.Date(2019, 10, 12, 0, 0, 0, 0, loc)
			Expect(epoch.EffectiveHoursInDay(t)).To(Equal(24))
		})

		It("should return 23 hours for days when PST switches to PDT", func() {
			loc, _ := time.LoadLocation("America/Los_Angeles")
			t := time.Date(2019, 3, 10, 12, 30, 0, 0, loc) // switch from PST to PDT happens at 2:00:00
			Expect(epoch.EffectiveHoursInDay(t)).To(Equal(23))
		})

		It("should return 25 hours for days when PDT switches to PST", func() {
			loc, _ := time.LoadLocation("America/Los_Angeles")
			t := time.Date(2019, 11, 3, 12, 30, 0, 0, loc) // switch from PDT to PST happens at 2:00:00
			Expect(epoch.EffectiveHoursInDay(t)).To(Equal(25))
		})
	})

	Context("TimeAddInterval", func() {
		When("interval with safe duration is given", func() {
			// it should be the same as `t.Add(i.Duration)`
			It("should add safe duration to time", func() {
				t := time.Now()
				i := epoch.MustParseInterval("5h")
				Expect(epoch.TimeAddInterval(t, i)).To(Equal(t.Add(i.Duration())))
			})
			It("should add safe duration to time with negative value", func() {
				t := time.Now()
				i := epoch.MustParseInterval("-5h")
				Expect(epoch.TimeAddInterval(t, i)).To(Equal(t.Add(i.Duration())))
			})
			It("should add safe duration to time with 0 value", func() {
				t := time.Now()
				i := epoch.MustParseInterval("0h")
				Expect(epoch.TimeAddInterval(t, i)).To(Equal(t.Add(i.Duration())))
			})

			It("should handle negative interval", func() {
				t := time.Date(2019, 10, 12, 5, 32, 0, 0, time.UTC)
				i := epoch.MustParseInterval("-1h")
				result := epoch.TimeAddInterval(t, i)

				Expect(epoch.TimeAddInterval(t, i)).To(Equal(t.Add(i.Duration())))
				Expect(result.String()).To(Equal("2019-10-12 04:32:00 +0000 UTC"))
			})

			It("should add interval with safe duration", func() {
				t := time.Date(2019, 10, 12, 5, 32, 0, 0, time.UTC)
				i := epoch.MustParseInterval("2h")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2019-10-12 07:32:00 +0000 UTC"))
			})
		})

		When("interval with non-safe duration is given", func() {
			It("should add interval", func() {
				t := time.Date(2019, 10, 12, 5, 32, 0, 0, time.UTC)
				i := epoch.MustParseInterval("2mo")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2019-12-12 05:32:00 +0000 UTC"))
			})
			It("should handle negative duration", func() {
				t := time.Date(2019, 10, 12, 5, 32, 0, 0, time.UTC)
				i := epoch.MustParseInterval("-2mo")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2019-08-12 05:32:00 +0000 UTC"))
			})
			It("should handle timezone", func() {
				loc, _ := time.LoadLocation("America/Los_Angeles")
				t := time.Date(2019, 10, 12, 5, 32, 0, 0, loc)
				i := epoch.MustParseInterval("2mo")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2019-12-12 05:32:00 -0800 PST"))
			})
			It("should handle daylight saving time", func() {
				loc, _ := time.LoadLocation("America/Los_Angeles")
				t := time.Date(2019, 10, 12, 5, 32, 0, 0, loc)
				i := epoch.MustParseInterval("2mo")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2019-12-12 05:32:00 -0800 PST"))
			})
			It("should handle negative", func() {
				a := time.Date(2019, 10, 12, 5, 32, 0, 0, time.UTC)
				i := epoch.MustParseInterval("-1mo")
				result := epoch.TimeAddInterval(a, i)
				Expect(result.String()).To(Equal("2019-09-12 05:32:00 +0000 UTC"))
			})
			It("should handle adding months that leads to next year", func() {
				t := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
				i := epoch.MustParseInterval("+2mo")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2020-02-15 00:00:00 +0000 UTC"))
			})

			It("should handle adding a month to a date at the end of the month (next month has fewer days)", func() {
				t := time.Date(2022, time.January, 31, 0, 0, 0, 0, time.UTC)
				i := epoch.MustParseInterval("1mo")
				result := epoch.TimeAddInterval(t, i)
				Expect(result.String()).To(Equal("2022-03-03 00:00:00 +0000 UTC"))
			})
		})
	})
})
