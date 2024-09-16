package pkg

import (
	"time"
)

type PDate struct {
	Year  int
	Month time.Month
	Day   int
}

func NewPDate(year int, month time.Month, day int) PDate {
	return PDate{Year: year, Month: month, Day: day}
}

type PTimeConfig struct {
	TimeZone string
}

type PTime struct {
	config *PTimeConfig
	now    time.Time
}

func NewPTime(cfg *PTimeConfig) *PTime {
	if cfg == nil {
		cfg = &PTimeConfig{
			TimeZone: "UTC",
		}
	}
	ptime := &PTime{
		config: cfg,
	}
	if cfg.TimeZone == "UTC" {
		ptime.now = time.Now().UTC()
	} else {
		ptime.now = time.Now().Local()
	}
	return ptime
}

func (pt *PTime) DateTime() time.Time {
	return pt.now
}
func (pt *PTime) Date() PDate {
	return NewPDate(pt.now.Date())
}
func (pt *PTime) DateString() string {
	return pt.now.Format(time.DateOnly)
}
func (pt *PTime) DateTimeString() string {
	return pt.now.Format(time.DateTime)
}
func (pt *PTime) TimeString() string {
	return pt.now.Format(time.TimeOnly)
}

func DateTimeFrom(datetime string) (*PTime, error) {
	d, err := time.Parse(time.DateTime, datetime)
	if err != nil {
		return nil, err
	}
	return &PTime{
		config: &PTimeConfig{
			TimeZone: "UTC",
		},
		now: d,
	}, nil
}

func DateFrom(date string) (*PTime, error) {
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return nil, err
	}
	return &PTime{
		config: &PTimeConfig{
			TimeZone: "UTC",
		},
		now: d,
	}, nil
}
