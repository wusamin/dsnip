package dsnip

import "time"

// GetDiff gets difference criteria and target.
func GetDiff(criteria time.Time, target time.Time) time.Duration {
	return criteria.Sub(target)
}

// GetDiffDay gets day difference criteria and target.
func GetDiffDay(criteria time.Time, target time.Time) time.Duration {
	return criteria.Sub(target) / (time.Hour * 24)
}

// GetDiffHour gets hour difference criteria and target.
func GetDiffHour(criteria time.Time, target time.Time) time.Duration {
	return criteria.Sub(target) / time.Hour
}

// GetDiffMinute gets minute difference criteria and target.
func GetDiffMinute(criteria time.Time, target time.Time) time.Duration {
	return criteria.Sub(target) / time.Minute
}
