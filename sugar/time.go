package sugar

import (
	"github.com/gocurr/good/consts"
	"time"
)

// ParseTime parses time-format string to time.Time
func ParseTime(s string, format ...string) time.Time {
	f := consts.DefaultTimeFormat
	if len(format) > 0 {
		f = format[0]
	}
	t, err := time.Parse(f, s)
	if err != nil {
		return time.Time{}
	}
	return t
}

// FormatTime formats time.Time to string
func FormatTime(t time.Time, format ...string) string {
	f := consts.DefaultTimeFormat
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}

// NowString formats time.Now to string
func NowString(format ...string) string {
	return FormatTime(time.Now(), format...)
}

// PointTime parses i to time.Time
func PointTime(ymd string, i int, d time.Duration) time.Time {
	day := ParseTime(ymd, consts.YMDFormat)
	return day.Add(time.Duration(i) * d)
}

// PointTimeString parses i to string
func PointTimeString(ymd string, i int, d time.Duration) string {
	t := PointTime(ymd, i, d)
	return FormatTime(t)
}
