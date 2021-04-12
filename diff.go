package dsnip

import "time"

// GetDiff gets difference criteria and target.
func GetDiff(from, to time.Time) time.Duration {
	return to.Sub(from)
}

// GetDiffDay gets day difference criteria and target.
func GetDiffDay(from, to time.Time) time.Duration {
	return to.Sub(from) / (time.Hour * 24)
}

// GetDiffHour gets hour difference criteria and target.
func GetDiffHour(from, to time.Time) time.Duration {
	return to.Sub(from) / time.Hour
}

// GetDiffMinute gets minute difference criteria and target.
func GetDiffMinute(from, to time.Time) time.Duration {
	return to.Sub(from) / time.Minute
}
