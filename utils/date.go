package utils

import (
	"fmt"
	"time"
)

func FormatDateInEnglish(t time.Time) string {
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
	weekdays := []string{
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}

	day := t.Day()
	month := months[t.Month()-1]
	year := t.Year()
	weekday := weekdays[t.Weekday()]

	return fmt.Sprintf("%s, %s %d, %d", weekday, month, day, year)
}
