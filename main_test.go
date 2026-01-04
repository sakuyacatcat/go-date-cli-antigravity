package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/sakuyacatcat/go-date-cli-antigravity/internal/view"
)

type MockTimeService struct {
	FixedTime time.Time
}

func (m *MockTimeService) Now() time.Time {
	return m.FixedTime
}

func (m *MockTimeService) Convert(t time.Time, location string) (time.Time, error) {
	if location == "Invalid/Location" {
		return time.Time{}, fmt.Errorf("unknown time zone %s", location)
	}
	loc, err := time.LoadLocation(location)
	if err != nil {
		// Fallback for test environment where some timezones might not be loaded if not using valid ones, 
		// but here we expect valid ones or manual handling. 
		// For simplicity in mock, let's just use real LoadLocation or simple offset if needed.
		// Actually, let's just use the Real service logic for conversion since that's not what limits us usually, 
		// OR better: just return time in UTC or fixed offset to be predictable without system files.
		return time.Time{}, err
	}
	return t.In(loc), nil
}

func TestRun(t *testing.T) {
	// Fixed time: 2026-01-01 12:00:00 UTC
	fixedTime := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	mockTimeSvc := &MockTimeService{FixedTime: fixedTime}
	realFormatter := &view.RealFormatter{}

	tests := []struct {
		name        string
		args        []string
		wantErr     bool
		wantOutPart string
	}{
		{
			name:        "Default execution",
			args:        []string{},
			wantErr:     false,
			wantOutPart: "Thu Jan  1 12:00:00 UTC 2026", // time.UnixDate format of fixedTime in UTC (since system time might be anything, but Now returns UTC here)
		},
		{
			name:        "With Timezone UTC",
			args:        []string{"-z", "UTC"},
			wantErr:     false,
			wantOutPart: "UTC",
		},
		{
			name:        "With ISO8601 Format",
			args:        []string{"-f", "ISO8601"},
			wantErr:     false,
			wantOutPart: "2026-01-01T12:00:00Z",
		},
		{
			name:        "With Long Flags",
			args:        []string{"--timezone", "UTC", "--format", "ISO8601"},
			wantErr:     false,
			wantOutPart: "Z",
		},
		{
			name:        "Invalid Timezone",
			args:        []string{"-z", "Invalid/Location"},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			// Ensure we reset mock if needed, but here it's stateless enough or shared is fine
			err := run(tt.args, &out, mockTimeSvc, realFormatter)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.wantOutPart != "" {
				if !strings.Contains(out.String(), tt.wantOutPart) {
					t.Errorf("run() output = %q, want to contain %q", out.String(), tt.wantOutPart)
				}
			}
		})
	}
}
