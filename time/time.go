package time

import "time"

const (
	DefaultFormat   = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
	YearMonthFormat = "2006-01"
)

// FormatDefaultTime
// convert time use format
func FormatDefaultTime(t time.Time, format string) string {
	if t.Unix() == 0 {
		t = time.Now()
	}
	if format == "" {
		format = DefaultFormat
	}
	return t.Format(format)
}

// TimestampToTime
// convert timestamp to time
func TimestampToTime(timestamp int64) time.Time {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	return time.Unix(timestamp, 0)
}

// TimestampToTimeStr
// convert timestamp to time string
func TimestampToTimeStr(timestamp int64, format string) string {
	if timestamp == 0 {
		return ""
	}
	if format == "" {
		format = DefaultFormat
	}
	return time.Unix(timestamp, 0).Format(format)
}
