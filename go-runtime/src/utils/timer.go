package utils

import "time"

func StartOfTomorow(today time.Time) time.Time {
	return today.UTC().Truncate(24 * time.Hour).Add(24 * time.Hour)
}

func StartOfToday(today time.Time) time.Time {
	return today.UTC().Truncate(24 * time.Hour)
}
