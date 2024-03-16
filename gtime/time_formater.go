package gtime

import (
	"strings"
	"time"
)

// All day gtime constants to layout.
const (
	LongMonthToken             = "January"
	MonthToken                 = "Jan"
	NumMonthToken              = "1"
	ZeroMonthToken             = "01"
	LongWeekDayToken           = "Monday"
	WeekDayToken               = "Mon"
	DayToken                   = "2"
	UnderDayToken              = "_2"
	ZeroDayToken               = "02"
	UnderYearDayToken          = "__2"
	ZeroYearDayToken           = "002"
	HourToken                  = "15"
	Hour12Token                = "3"
	ZeroHour12Token            = "03"
	MinuteToken                = "4"
	ZeroMinuteToken            = "04"
	SecondToken                = "5"
	ZeroSecondToken            = "05"
	LongYearToken              = "2006"
	YearToken                  = "06"
	PMToken                    = "PM"
	pmToken                    = "pm"
	TZToken                    = "MST"
	ISO8601TZToken             = "Z0700" // prints Z for UTC
	ISO8601SecondsTZToken      = "Z070000"
	ISO8601ShortTZToken        = "Z07"
	ISO8601ColonTZToken        = "Z07:00" // prints Z for UTC
	ISO8601ColonSecondsTZToken = "Z07:00:00"
	NumTZToken                 = "-0700" // always numeric
	NumSecondsTzToken          = "-070000"
	NumShortTZToken            = "-07"    // always numeric
	NumColonTZToken            = "-07:00" // always numeric
	NumColonSecondsTZToken     = "-07:00:00"
	FracSecond0Token           = ".0"
	FracSecond9Token           = ".9"
	YYMMDDHHMMMM               = "2006-01-02T15:04:05Z"
)

// Formatter interface defines all methods that can be chained.
type Formatter interface {
	AppendToken(t string) Formatter
	GenerateMask() Formatter
	GenerateWithAllBetween(token string) Formatter
	Convert() string
}

// MakeFormatter builds an object with the gtime to be formatted.
func MakeFormatter(t time.Time) FormattedDate {
	return FormattedDate{timeToConvert: t}
}

// FormattedDate main object structure .
type FormattedDate struct {
	maskSlice     []string
	maskDone      string
	timeToConvert time.Time
}

// AppendToken appends a token to the mask.
func (f *FormattedDate) AppendToken(t string) Formatter {
	f.maskSlice = append(f.maskSlice, t)
	return f
}

// GenerateMask builds a mask.
func (f *FormattedDate) GenerateMask() Formatter {
	f.maskDone = strings.Join(f.maskSlice, "")
	return f
}

// GenerateWithAllBetween builds a mask joining all parts with the given string.
func (f *FormattedDate) GenerateWithAllBetween(token string) Formatter {
	f.maskDone = strings.Join(f.maskSlice, token)
	return f
}

// Convert do the formatting conversion.
func (f *FormattedDate) Convert() string {
	mask := f.GetMask()
	return f.timeToConvert.Format(mask)
}

// GetMask returns the mask.
func (f *FormattedDate) GetMask() string {
	if len(f.maskDone) == 0 {
		f.GenerateMask()
	}
	return f.maskDone
}
