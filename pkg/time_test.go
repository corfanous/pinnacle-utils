package pkg_test

import (
	"testing"
	"time"

	. "github.com/corfanous/pinnacle-utils/pkg"
)

func TestNewPTime(t *testing.T) {
	if NewPTime(nil).DateString() == time.Now().Format(time.DateOnly) {
		t.Logf("NewPTime(nil) successfully created %v", *NewPTime(nil))
	} else {
		t.Error("NewPTime(nil)): test failure")
	}
}
func TestDateString(t *testing.T) {
	ptime := NewPTime(nil)
	if ptime.DateString() == time.Now().Format(time.DateOnly) {
		t.Logf("ptime.DateString() is: %s\n", time.Now().Format(time.DateOnly))
	} else {
		t.Error("ptime.DateString(): test failure")
	}
}
func TestDateTimeString(t *testing.T) {
	ptime := NewPTime(nil)
	if ptime.DateTimeString() == ptime.DateTime().Format(time.DateTime) {
		t.Logf("ptime.DateTimeString() is: %s\n", ptime.DateTimeString())
	} else {
		t.Error("invalid execution")
	}
}

func TestDateFrom(t *testing.T) {
	suits := []struct {
		Name    string
		Value   string
		Expects bool
	}{
		{
			Name:    "Date only",
			Value:   time.Now().Format(time.DateOnly),
			Expects: true,
		},
		{
			Name:    "Time only",
			Value:   time.Now().Format(time.TimeOnly),
			Expects: false,
		},
		{
			Name:    "Date with Time",
			Value:   time.Now().Format(time.DateTime),
			Expects: false,
		},
	}
	for _, suit := range suits {
		_, err := DateFrom(suit.Value)
		if result := err == nil; result == suit.Expects {
			t.Logf("%s: %s", suit.Name, suit.Value)
		} else {
			t.Errorf("%s: %v\n", suit.Name, err)
		}
	}
}

func TestDateTimeFrom(t *testing.T) {
	suits := []struct {
		Name    string
		Value   string
		Expects bool
	}{
		{
			Name:    "Date only",
			Value:   time.Now().Format(time.DateOnly),
			Expects: false,
		},
		{
			Name:    "Time only",
			Value:   time.Now().Format(time.TimeOnly),
			Expects: false,
		},
		{
			Name:    "Date with Time",
			Value:   time.Now().Format(time.DateTime),
			Expects: true,
		},
	}
	for _, suit := range suits {
		_, err := DateTimeFrom(suit.Value)
		if result := err == nil; result == suit.Expects {
			t.Logf("%s: %s", suit.Name, suit.Value)
		} else {
			t.Errorf("%s: %v\n", suit.Name, err)
		}
	}
}
