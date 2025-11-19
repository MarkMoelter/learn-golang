package shapes_test

import (
	"testing"

	"github.com/MarkMoelter/learn-golang/shapes"
)

func TestPerimeter(t *testing.T) {
	rectangle := shapes.Rectangle{10.0, 10.0}
	got := shapes.Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name 	string
		shape 	shapes.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: shapes.Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: shapes.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: shapes.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
