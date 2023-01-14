package epoch_test

import (
	"errors"
	"github.com/aahainc/epoch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Intervals", func() {
	Context("ParseInterval", func() {
		DescribeTable("valid input is given", func(inputStr string, expectedVal float64, expectedUnit epoch.Unit) {
			interval, err := epoch.ParseInterval(inputStr)
			Expect(err).Should(Succeed())
			Expect(interval).NotTo(BeNil())
			Expect(interval.Value).To(Equal(expectedVal))
			Expect(interval.Unit).To(Equal(expectedUnit))
		},
			Entry("5 seconds", "5s", 5.0, epoch.UnitSecond),
			Entry("3 minutes", "3m", 3.0, epoch.UnitMinute),
			Entry("7 hours", "7h", 7.0, epoch.UnitHour),
			Entry("2 days", "2d", 2.0, epoch.UnitDay),
			Entry("4 weeks", "4w", 4.0, epoch.UnitWeek),
			Entry("6 months", "6mo", 6.0, epoch.UnitMonth),
			Entry("10 years", "10y", 10.0, epoch.UnitYear),
			Entry("1.5 seconds", "1.5s", 1.5, epoch.UnitSecond),
			Entry("3.14 minutes", "3.14m", 3.14, epoch.UnitMinute),
			Entry("6.28 hours", "6.28h", 6.28, epoch.UnitHour),
			Entry("1.41 days", "1.41d", 1.41, epoch.UnitDay),
			Entry("3.33 weeks", "3.33w", 3.33, epoch.UnitWeek),
			Entry("5.55 months", "5.55mo", 5.55, epoch.UnitMonth),
			Entry("7.77 years", "7.77y", 7.77, epoch.UnitYear),

			Entry("-5 minutes", "-5m", -5.0, epoch.UnitMinute),
			Entry("0 second", "0s", 0.0, epoch.UnitSecond),
			Entry("-0.5 hour", "-0.5h", -0.5, epoch.UnitHour),
			Entry("0 day", "0d", 0.0, epoch.UnitDay),
			Entry("-2 week", "-2w", -2.0, epoch.UnitWeek),
			Entry("-3 year", "-3y", -3.0, epoch.UnitYear),
		)

		DescribeTable("invalid input is given", func(inputStr string, expectedError error) {
			interval, err := epoch.ParseInterval(inputStr)
			Expect(interval).To(BeNil())
			Expect(errors.Is(err, expectedError)).To(BeTrue(), inputStr)
		},
			Entry("empty input", "", epoch.ErrInvalidFormat),
			Entry("invalid unit", "5x", epoch.ErrInvalidUnit),

			// todo: it's actually ignores `.3`, should fail on InvalidFormat
			//Entry("invalid value", "5.5.3m", epoch.ErrInvalidFormat),
		)

	})

	Context("MustParseInterval", func() {
		It("parses a valid interval without error", func() {
			interval := epoch.MustParseInterval("5s")
			Expect(interval.Value).To(Equal(5.0))
			Expect(interval.Unit).To(Equal(epoch.UnitSecond))
		})
		It("panics on an invalid interval", func() {
			Expect(func() { epoch.MustParseInterval("invalid") }).To(Panic())
		})
	})

	Context("interval.Duration()", func() {
		DescribeTable("valid input is given", func(input epoch.Interval, expectedDuration time.Duration) {
			duration := input.Duration()
			Expect(duration).To(Equal(expectedDuration), input.String())
		},
			Entry("5 seconds", epoch.Interval{5, epoch.UnitSecond}, 5*time.Second),
			Entry("5 minutes", epoch.Interval{5, epoch.UnitMinute}, 5*time.Minute),
			Entry("5 hours", epoch.Interval{5, epoch.UnitHour}, 5*time.Hour),
			Entry("5 days", epoch.Interval{5, epoch.UnitDay}, 5*24*time.Hour),
			Entry("5 weeks", epoch.Interval{5, epoch.UnitWeek}, 5*7*24*time.Hour),
			Entry("5.5 seconds", epoch.Interval{5.5, epoch.UnitSecond}, 5500*time.Millisecond),
			Entry("0 second", epoch.Interval{0, epoch.UnitSecond}, 0*time.Second),
			Entry("-5 seconds", epoch.Interval{-5, epoch.UnitSecond}, -5*time.Second),
		)

		DescribeTable("invalid input is given", func(input epoch.Interval) {
			Expect(func() { input.Duration() }).To(Panic())
		},
			Entry("5 months", epoch.Interval{5, epoch.UnitMonth}),
			Entry("5 years", epoch.Interval{5, epoch.UnitYear}),
			Entry("-5 months", epoch.Interval{-5, epoch.UnitMonth}),
			Entry("-5 years", epoch.Interval{-5, epoch.UnitYear}),
		)
	})

	Context("IsSafeDuration()", func() {
		It("returns true for safe durations", func() {
			intervals := []epoch.Interval{
				{5, epoch.UnitSecond},
				{5, epoch.UnitMinute},
				{5, epoch.UnitHour},
				{5, epoch.UnitDay},
				{5, epoch.UnitWeek},
			}

			for _, interval := range intervals {
				Expect(interval.IsSafeDuration()).To(BeTrue())
			}
		})

		It("returns false for unsafe durations", func() {
			intervals := []epoch.Interval{
				{5, epoch.UnitMonth},
				{5, epoch.UnitYear},
			}

			for _, interval := range intervals {
				Expect(interval.IsSafeDuration()).To(BeFalse())
			}
		})
	})

	Context("ExtractDateParts", func() {
		It("should extract date parts for day unit", func() {
			i := epoch.Interval{Value: 2, Unit: epoch.UnitDay}
			y, m, d := i.ExtractDateParts()
			Expect(y).To(Equal(0))
			Expect(m).To(Equal(0))
			Expect(d).To(Equal(2))
		})

		It("should extract date parts for week unit", func() {
			i := epoch.Interval{Value: 2, Unit: epoch.UnitWeek}
			y, m, d := i.ExtractDateParts()
			Expect(y).To(Equal(0))
			Expect(m).To(Equal(0))
			Expect(d).To(Equal(14))
		})

		It("should extract date parts for month unit", func() {
			i := epoch.Interval{Value: 2, Unit: epoch.UnitMonth}
			y, m, d := i.ExtractDateParts()
			Expect(y).To(Equal(0))
			Expect(m).To(Equal(2))
			Expect(d).To(Equal(0))
		})

		It("should extract date parts for year unit", func() {
			i := epoch.Interval{Value: 2, Unit: epoch.UnitYear}
			y, m, d := i.ExtractDateParts()
			Expect(y).To(Equal(2))
			Expect(m).To(Equal(0))
			Expect(d).To(Equal(0))
		})

		It("should extract date parts for non-date unit", func() {
			i := epoch.Interval{Value: 2, Unit: epoch.UnitSecond}
			y, m, d := i.ExtractDateParts()
			Expect(y).To(Equal(0))
			Expect(m).To(Equal(0))
			Expect(d).To(Equal(0))
		})
	})
})
