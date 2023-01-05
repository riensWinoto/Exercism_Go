package main

import "fmt"

type Clock struct {
	minute int
}

func New(h, m int) Clock {
	tempMin := m + h*60
	tempMin = tempMin % (24 * 60)
	if tempMin < 0 {
		tempMin = tempMin + (24 * 60)
	}
	return Clock{minute: tempMin}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}
