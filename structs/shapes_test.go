package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	expected := rectangle.Perimeter()
	want := 40.0
	if expected != want {
		t.Errorf("%.2f expected perimeter of %.2f ", expected, want)
	}
}
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {

		expected := tt.shape.Area()

		if expected != tt.hasArea {
			t.Errorf("%#v got %g expected area of %g ", tt.shape, expected, tt.hasArea)
		}
	}

}
