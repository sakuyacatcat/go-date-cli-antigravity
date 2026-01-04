package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sakuyacatcat/go-date-cli-antigravity/internal/infrastructure/clock"
	"github.com/sakuyacatcat/go-date-cli-antigravity/internal/service"
	"github.com/sakuyacatcat/go-date-cli-antigravity/internal/view"
)

func main() {
	realClock := &clock.RealClock{}
	timeSvc := &service.RealTimeService{Clock: realClock}
	formatter := &view.RealFormatter{}

	if err := run(os.Args[1:], os.Stdout, timeSvc, formatter); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer, timeSvc service.TimeService, formatter view.Formatter) error {
	fs := flag.NewFlagSet("date-cli", flag.ContinueOnError)
	
	// Flags
	var (
		tzShort, tzLong string
		fmtShort, fmtLong string
	)

	// Since flag package doesn't support direct aliasing with same variable cleanly without potential issues if defaults differ (here defaults are empty),
	// we CAN bind to same variable effectively if we are careful.
	// But standard pattern often uses separate vars or manual handling.
	// Let's use separate vars and resolve.
	
	fs.StringVar(&tzShort, "z", "", "Timezone (e.g. UTC, Asia/Tokyo)")
	fs.StringVar(&tzLong, "timezone", "", "Timezone (e.g. UTC, Asia/Tokyo)")
	
	fs.StringVar(&fmtShort, "f", "", "Output format (e.g. ISO8601)")
	fs.StringVar(&fmtLong, "format", "", "Output format (e.g. ISO8601)")

	if err := fs.Parse(args); err != nil {
		return err
	}

	// Resolve flags
	tz := tzShort
	if tzLong != "" {
		tz = tzLong
	}

	format := fmtShort
	if fmtLong != "" {
		format = fmtLong
	}

	// Logic
	t := timeSvc.Now()
	
	if tz != "" {
		var err error
		t, err = timeSvc.Convert(t, tz)
		if err != nil {
			return fmt.Errorf("invalid timezone: %w", err)
		}
	}

	s := formatter.Format(t, format)
	fmt.Fprintln(out, s)

	return nil
}
