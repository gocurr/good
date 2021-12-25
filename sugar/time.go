package sugar

import (
	"github.com/gocurr/good/consts"
	"time"
)

// ParseTime returns time.Time which is parsed from the give string.
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

// FormatTime return a new string which is parsed from the given time t.
func FormatTime(t time.Time, format ...string) string {
	f := consts.DefaultTimeFormat
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}

// NowString returns a new string which is formatted from time.Now.
func NowString(format ...string) string {
	return FormatTime(time.Now(), format...)
}

// PointTime returns time.Time which is parsed by the given ymd, point: i, and duration: d.
func PointTime(ymd string, i int, d time.Duration) time.Time {
	day := ParseTime(ymd, consts.YMDFormat)
	return day.Add(time.Duration(i) * d)
}

// PointTimeString returns a new string which is parsed by the given ymd, point: i, and duration: d.
func PointTimeString(ymd string, i int, d time.Duration) string {
	t := PointTime(ymd, i, d)
	return FormatTime(t)
}
