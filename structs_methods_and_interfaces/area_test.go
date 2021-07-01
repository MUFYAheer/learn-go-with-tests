package structsmethodsandinterfaces

import (
	"testing"
)

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{4, 5}, 20.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Circle of radius 5", Circle{5}, 78.53981633974483},
		{"Triangle", Triangle{4, 5}, 10},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
