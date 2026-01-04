package clock

import "time"

// Clock abstracts the time package functions.
type Clock interface {
	Now() time.Time
	LoadLocation(name string) (*time.Location, error)
}

// RealClock is the concrete implementation that uses the time package.
type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}

func (c *RealClock) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
