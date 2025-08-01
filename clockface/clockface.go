package clockface

import (
	"math"
	"time"
)

const (
	clockCenterX = 150
	clockCenterY = 150
	hourHandLen = 50
	minHandLen	= 80
	secHandLen	= 90
)

type Point struct {
	X float64
	Y float64
}

func SecondsInRadians(t time.Time) float64 {
	return (float64(t.Second()) / 30) * math.Pi
}

func SecondHandPoint(t time.Time) Point {
	rad := SecondsInRadians(t)
	return Point{X: math.Sin(rad), Y: math.Cos(rad)}
}
