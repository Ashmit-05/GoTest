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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		assertCorrectMessage(t, got, want)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{7.0}
		got := circle.Area()
		want := 153.93804002589985

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %g got %g", want, got)
	}
}
