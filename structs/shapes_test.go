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
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		expected := shape.Area()
		if expected != want {
			t.Errorf("%g expected area of %g ", expected, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)

	})

}
