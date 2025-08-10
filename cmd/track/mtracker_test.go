package main

import (
	"testing"
	"time"
)

func TestFormatMouseEvent(t *testing.T) {
	ts := time.Date(2030, 4, 5, 15, 30, 12, 0, time.UTC)
	x, y := 123, 45
	expected := "2030-04-05T15:30:12Z mouse => 123, 45\n"
	got := formatMouseEvent(ts, x, y)
	if got != expected {
		t.Errorf("unexpected format. got %q want %q", got, expected)
	}

	// edge: negative and large values
	ts2 := time.Date(9999, 1, 1, 1, 1, 1, 0, time.UTC)
	got2 := formatMouseEvent(ts2, -99999, 1000000)
	exp2 := "9999-01-01T01:01:01Z mouse => -99999, 1000000\n"
	if got2 != exp2 {
		t.Errorf("unexpected format2. got %q want %q", got2, exp2)
	}
}

func TestStartupMessage(t *testing.T) {
	exp := "--- ctrl + q to quit / ctrl + shift + hold to record position ---"
	if startupMessage() != exp {
		t.Errorf("wrong startup message: %q", startupMessage())
	}
}
