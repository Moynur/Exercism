package clock

import "fmt"

// Clock basic clock
type Clock struct {
	hour   int
	minute int
}

// New Clock
func New(hour, minute int) Clock {
	var c Clock
	c.minute = minute +hour*60
	return c
}

// Add comment
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c
}

// Subtract comment
func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	return c
}

// String comment
func (c Clock) String() string {
	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}
	for c.hour < 0 {
		c.hour += 24
	}

	c.hour += int(c.minute / 60)
	c.minute = c.minute % 60
	c.hour = c.hour % 24
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
