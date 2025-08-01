package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondsInRadians(t time.Time) float64 {
	return (float64(t.Second()) / 30) * math.Pi
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) + (float64(t.Minute()) / 30) * math.Pi
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
