package clockface

import (
	"math"
	"time"
)

const (
	secondsInClock = 60
	minutesInClock = 60
	hoursInClock = 12
)

type Point struct {
	X float64
	Y float64
}

func SecondsInRadians(t time.Time) float64 {
	return (float64(t.Second()) / secondsInClock) * (2 * math.Pi)
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinutesInRadians(t time.Time) float64 {
	offset := SecondsInRadians(t) / minutesInClock
	currentMinute := (float64(t.Minute()) / minutesInClock) * (2 * math.Pi)
	return offset + currentMinute
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func HoursInRadians(t time.Time) float64 {
	offset := MinutesInRadians(t) / hoursInClock
	currentHour := ((float64(t.Hour() % hoursInClock)) / hoursInClock) * (2 * math.Pi)
	return offset + currentHour
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
