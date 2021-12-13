package area

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, area: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, area: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()

			if got != test.area {
				t.Errorf("%#v, got %v, want %v", test.shape, got, test.area)
			}
		})
	}
}

// go test -run TestArea/Circle
