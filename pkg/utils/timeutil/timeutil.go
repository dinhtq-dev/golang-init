package timeutil

import (
	"time"
)

// GetCurrentDate returns the current date in YYYY-MM-DD format
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// ParseDate converts a string to a time object based on a specific format
func ParseDate(dateStr string, format string) (time.Time, error) {
	return time.Parse(format, dateStr)
}

// FormatDate formats a time object to a string based on a specific format
func FormatDate(t time.Time, format string) string {
	return t.Format(format)
}

// AddDays adds a specified number of days to a date
func AddDays(t time.Time, days int) time.Time {
	return t.Add(time.Duration(days) * time.Hour * 24)
}

// DaysBetween calculates the number of days between two dates
func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}
