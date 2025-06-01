package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	shape := Rectangle{1.2, 2.0}
	got := Perimeter(shape)
	want := 6.4
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	shape := Rectangle{1.0, 1.2}
	for i := 0; i < b.N; i++ {
		Perimeter(shape)
	}
}

func BenchmarkArea(b *testing.B) {
	shape := Triangle{Base: 10, Height: 2}
	for i := 0; i < b.N; i++ {
		shape.Area()
	}
}

func ExamplePerimeter() {
	rectangle := Rectangle{1.2, 2.0}
	perimeter := Perimeter(rectangle)
	fmt.Println(perimeter)
	// Output: 6.4
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea
			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})
	}
}
