package service

import (
	"testing"
	"time"
)

type MockClock struct {
	FixedTime time.Time
}

func (m *MockClock) Now() time.Time {
	return m.FixedTime
}

func (m *MockClock) LoadLocation(name string) (*time.Location, error) {
	// For testing purposes, we can return nil or error based on name.
	// Since we can't easily synthesize a Location without loading from system or creating it effectively,
	// and Time.In(loc) needs a valid loc.
	// However, for unit testing Convert, we want to verify it calls LoadLocation and uses the result.
	// If name is "Invalid/Location", return error.
	if name == "Invalid/Location" {
		_, err := time.LoadLocation(name)
		return nil, err
	}
	// For valid cases, in a true unit test we might want to mock Location too but Location is hard to mock directly as struct.
	// So we can fallback to real LoadLocation for "valid" inputs if we assume test env has them,
	// OR we can return time.UTC for everything if we just want to test flow, but ensuring correct conversion math requires real location data.
	// Let's use real LoadLocation for valid string to strictly test math, but via the interface.
	return time.LoadLocation(name)
}

func TestRealTimeService_Now(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock := &MockClock{FixedTime: fixedTime}
	s := &RealTimeService{Clock: mockClock}
	
	got := s.Now()
	if !got.Equal(fixedTime) {
		t.Errorf("Now() = %v, want %v", got, fixedTime)
	}
}

func TestRealTimeService_Convert(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	mockClock := &MockClock{FixedTime: fixedTime}
	s := &RealTimeService{Clock: mockClock}
	
	// Create a stable time for testing: 2023-01-01 12:00:00 UTC
	baseTime := fixedTime

	tests := []struct {
		name     string
		input    time.Time
		location string
		wantLoc  string
		wantErr  bool
	}{
		{
			name:     "Convert to UTC",
			input:    baseTime,
			location: "UTC",
			wantLoc:  "UTC",
			wantErr:  false,
		},
		{
			name:     "Convert to JST (Asia/Tokyo)",
			input:    baseTime,
			location: "Asia/Tokyo",
			wantLoc:  "Asia/Tokyo",
			wantErr:  false,
		},
		{
			name:     "Convert to Invalid Location",
			input:    baseTime,
			location: "Invalid/Location",
			wantLoc:  "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Convert(tt.input, tt.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Location().String() != tt.wantLoc {
					t.Errorf("Convert() location = %v, want %v", got.Location().String(), tt.wantLoc)
				}
				// Verify the instant in time is the same
				if !got.Equal(tt.input) {
					t.Errorf("Convert() time instant changed: got %v, input %v", got, tt.input)
				}
			}
		})
	}
}
