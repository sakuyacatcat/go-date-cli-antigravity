package view

import "time"

// Formatter defines the interface for formatting time.
type Formatter interface {
	Format(t time.Time, pattern string) string
}

// RealFormatter is the concrete implementation of Formatter.
type RealFormatter struct{}

// Format formats the time according to the pattern.
func (f *RealFormatter) Format(t time.Time, pattern string) string {
	if pattern == "" {
		return t.Format(time.UnixDate)
	}
	if pattern == "ISO8601" {
		return t.Format(time.RFC3339)
	}
	return t.Format(pattern)
}
