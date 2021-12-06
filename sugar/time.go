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
