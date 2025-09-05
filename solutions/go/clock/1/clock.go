package clock

import "fmt"

// Clock represents the time in minutes since midnight.
type Clock struct {
	minutes int
}

// minutesInDay is the total number of minutes in a day (24*60)
const minutesInDay = 24 * 60

// normalize ensures minutes is within 0 <= minutes < 1440
func normalize(m int) int {
	m = m % minutesInDay
	if m < 0 {
		m += minutesInDay
	}
	return m
}

// New returns a new Clock value set to h:m.
func New(h, m int) Clock {
	total := h*60 + m
	return Clock{normalize(total)}
}

// Add returns a new Clock, minutes added.
func (c Clock) Add(m int) Clock {
	return Clock{normalize(c.minutes + m)}
}

// Subtract returns a new Clock, minutes subtracted.
func (c Clock) Subtract(m int) Clock {
	return Clock{normalize(c.minutes - m)}
}

// String returns the time in "HH:MM" format.
func (c Clock) String() string {
	h := c.minutes / 60
	m := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
