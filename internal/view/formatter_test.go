package view

import (
	"testing"
	"time"
)

func TestRealFormatter_Format(t *testing.T) {
	s := &RealFormatter{}
	
	// 2023-01-01 12:34:56 UTC
	testTime := time.Date(2023, 1, 1, 12, 34, 56, 0, time.UTC)

	tests := []struct {
		name    string
		time    time.Time
		pattern string
		want    string
	}{
		{
			name:    "Custom Pattern",
			time:    testTime,
			pattern: "2006-01-02 15:04:05",
			want:    "2023-01-01 12:34:56",
		},
		{
			name:    "ISO8601 Pattern",
			time:    testTime,
			pattern: "ISO8601",
			want:    "2023-01-01T12:34:56Z",
		},
		{
			name:    "Default UnixDate (Empty pattern for testing purpose, though CLI has default)",
			time:    testTime,
			pattern: "", // If empty, what should happen? Or is default handled by caller?
                            // Let's assume Formatter just blindly formats if possible, or we define a behavior.
                            // The plan said "Default format: time.UnixDate".
                            // Let's make empty pattern fallback to UnixDate for convenience.
			want:    testTime.Format(time.UnixDate),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Format(tt.time, tt.pattern)
			if got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
