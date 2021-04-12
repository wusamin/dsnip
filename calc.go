package dsnip

import "time"

// AddDay adds days to time.Time.
func AddDay(t time.Time, day int) time.Time {
	return t.Add(time.Duration(day) * (time.Hour * 24))
}

// AddHour adds hour to time.Time.
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Duration(hour) * time.Hour)
}

// AddMinute adds minutes to time.Time.
func AddMinute(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddSeconds adds seconds to time.Time.
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}
