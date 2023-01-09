package main

import "time"

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
