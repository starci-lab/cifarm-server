package utils

import "time"

func StartOfNextDay(current time.Time) time.Time {
	return current.UTC().Truncate(24 * time.Hour).Add(24 * time.Hour)
}
