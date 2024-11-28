package main 

import "testing"

type Rectangle struct {
	Width float64
	Height float64 
}

type Circle struct {
	Radius float64
}

func TestPerimetre(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimetre(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
  
checkArea := func(t testing.TB, shape Shape, want float64)

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5}
		got := circle.Area()
		want := 78.53981633974483

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}