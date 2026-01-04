package service

import (
	"time"

	"github.com/sakuyacatcat/go-date-cli-antigravity/internal/infrastructure/clock"
)

// TimeService defines operations for retrieving and converting time.
type TimeService interface {
	Now() time.Time
	Convert(t time.Time, location string) (time.Time, error)
}

// RealTimeService is the concrete implementation of TimeService.
type RealTimeService struct {
	Clock clock.Clock
}

// Now returns the current local time.
func (s *RealTimeService) Now() time.Time {
	return s.Clock.Now()
}

// Convert converts the given time to the specified location.
func (s *RealTimeService) Convert(t time.Time, location string) (time.Time, error) {
	loc, err := s.Clock.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}
