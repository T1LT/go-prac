// Package gigasecond provides functions to determine the date after a "gigasecond".
package gigasecond

import "time"

// AddGigasecond takes in a time and adds one gigasecond to it
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
