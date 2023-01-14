package epoch

import (
	"time"
)

// TruncateToHour truncates the given time to the hour by rounding down
func TruncateToHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func TruncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// RoundUpToHour rounds the given time up to the nearest hour (rounding right)
func RoundUpToHour(t time.Time) time.Time {
	return TruncateToHour(t).Add(1 * time.Hour)
}

// RoundUpToDay rounds the given time up to the nearest day (rounding right)
func RoundUpToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
}

func EffectiveHoursInDay(t time.Time) int {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endOfDay := RoundUpToDay(t)

	duration := endOfDay.Sub(startOfDay)
	return int(duration.Hours())
}

// TimeAddInterval adds the given interval to the given time and returns the resulting time.
// If the interval is a safe duration (can be converted to a precise time.Duration),
// it will use the t.Add(i.Duration) method.
// Otherwise, for intervals based on months and years, it will use t.AddDate(i.ExtractDateParts())
func TimeAddInterval(t time.Time, i *Interval) time.Time {
	if i.IsSafeDuration() {
		return t.Add(i.Duration())
	}

	return t.AddDate(i.ExtractDateParts())
}
