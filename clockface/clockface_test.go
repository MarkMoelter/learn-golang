package clockface_test

import (
	"math"
	"testing"
	"time"

	"github.com/MarkMoelter/learn-golang/clockface"
)

func TestSecondHand(t *testing.T) {
	t.Run("test seconds in radians", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(0, 0, 30), math.Pi},
			{simpleTime(0, 0, 0), 0},
			{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
			{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := clockface.SecondsInRadians(c.time)
				if !roughlyEqualFloat64(got, c.angle) {
					t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
				}
			})
		}
	})
	t.Run("test second hand point", func(t *testing.T) {
		cases := []struct {
			time time.Time
			point clockface.Point
		}{
			{simpleTime(0, 0, 30), clockface.Point{0, -1}},
			{simpleTime(0, 0, 45), clockface.Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := clockface.SecondHandPoint(c.time)
				if !roughlyEqualPoint(got, c.point) {
					t.Fatalf("Wanted %v Point, but got %v", c.point, got)
				}
			})
		}
	})
}

func TestMinuteHand(t *testing.T) {
	t.Run("test minutes in radians", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(0, 30, 0), math.Pi},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := clockface.MinutesInRadians(c.time)
				if got != c.angle {
					t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
				}
			})
		}
	})
	t.Run("test minute hand point", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			point clockface.Point
		}{
			{simpleTime(0, 30, 0), clockface.Point{0, -1}},
			{simpleTime(0, 45, 0), clockface.Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := clockface.MinuteHandPoint(c.time)
				if !roughlyEqualPoint(got, c.point) {
					t.Fatalf("Wanted %v Point, but got %v", c.point, got)
				}
			})
		}
	})
}

func TestHourHand(t *testing.T) {
	t.Run("test hours in radians", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(6, 0, 0), math.Pi},
			{simpleTime(0, 0, 0), 0},
			{simpleTime(21, 0, 0), math.Pi * 1.5},
			{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},

		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := clockface.HoursInRadians(c.time)
				if !roughlyEqualFloat64(got, c.angle) {
					t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
				}
			})
		}
	})
	t.Run("test hour hand point", func(t *testing.T)  {
		cases := []struct {
			time  time.Time
			point clockface.Point
		}{
			{simpleTime(6, 0, 0), clockface.Point{0, -1}},
			{simpleTime(21, 0, 0), clockface.Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := clockface.HourHandPoint(c.time)
				if !roughlyEqualPoint(got, c.point) {
					t.Fatalf("Wanted %v Point, but got %v", c.point, got)
				}
			})
		}
	})
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)	
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
