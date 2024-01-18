package structsmethodsinterface

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	assertCorrectMessage(t, got, want)

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// {Rectangle{10.0, 10.0}, 100.0},
		// {Circle{7.0}, 153.93804002589985},
		// {Triangle{15.0, 9.3}, 69.75},
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{7.0}, hasArea: 153.93804002589985},
		{name: "Triangle", shape: Triangle{15.0, 9.3}, hasArea: 69.75},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea

			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})
	}

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	//
	// 	assertCorrectMessage(t, got, want)
	//
	// }
	//
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkArea(t, rectangle, 100.0)
	//
	// })
	//
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{7.0}
	// 	checkArea(t, circle, 153.93804002589985)
	// })

}

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %g got %g", want, got)
	}
}
