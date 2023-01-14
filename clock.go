package epoch

import "time"

type Clock interface {
	Now() time.Time
}

var _ Clock = &DefaultClock{}

type DefaultClock struct{}

func (c *DefaultClock) Now() time.Time {
	return time.Now()
}

func NewDefaultClock() *DefaultClock {
	return &DefaultClock{}
}

type StaticClock struct {
	fixedTime time.Time
}

var _ Clock = &StaticClock{}

func (c *StaticClock) Now() time.Time {
	return c.fixedTime
}

func NewStaticClock(fixedTime time.Time) *StaticClock {
	return &StaticClock{fixedTime: fixedTime}
}
