package clock

import (
	"log"
	"math"
	"strconv"
)

const testVersion = 4

type Clock struct {
	h int
	m int
}

func New(hour, minute int) Clock {
	return Clock{h: hour, m: minute}
}

func (c Clock) String() string {
	totalMinutes := c.h*60 + c.m
	hour := int(math.Floor(float64(totalMinutes)/60)) % 24
	if hour < 0 {
		hour += 24
	}
	minute := totalMinutes % 60
	if minute < 0 {
		minute += 60
	}

	log.Printf("total minutes: %d\n", totalMinutes)
	log.Printf("hour: %d, minute: %d\n", hour, minute)
	rep := ""
	if hour < 10 {
		rep += "0"
	}
	rep += strconv.FormatInt(int64(hour), 10)
	rep += ":"

	if minute < 10 {
		rep += "0"
	}
	rep += strconv.FormatInt(int64(minute), 10)
	log.Printf("%d, %d, %s\n", hour, minute, rep)
	return rep
}

func (c Clock) Add(minutes int) Clock {
	totalMinutes := c.h*60 + c.m + minutes

	hour := int(math.Floor(float64(totalMinutes)/60)) % 24
	if hour < 0 {
		hour += 24
	}
	minute := totalMinutes % 60
	if minute < 0 {
		minute += 60
	}
	return Clock{h: hour, m: minute}
}
