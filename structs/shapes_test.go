package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	var (
		want      = 40.0
		rectangle = Rectangle{10.0, 10.0}
		got       = Perimeter(rectangle)
	)
	if want != got {
		t.Errorf("want: %.2f, got: %.2f", want, got)
	}
}

// ----------------

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 31.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if tt.want != got {
				t.Errorf("shape: %#v, want: %g, got %g", tt.shape, tt.want, got)
			}
		})
	}
}
