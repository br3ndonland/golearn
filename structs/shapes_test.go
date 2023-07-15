package main

import "testing"

func assertCorrectFloat(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	assertCorrectFloat(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		assertCorrectFloat(t, got, want)
	}
	t.Run("rectangle", func(t *testing.T) {
		shape := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, shape, want)
	})
	t.Run("circle", func(t *testing.T) {
		shape := Circle{10}
		want := 314.1592653589793
		checkArea(t, shape, want)
	})
}
