package structsmethodsinterface

import "math"

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth

}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
  return 0.5 * t.Base * t.Height
}

// These are functions
// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Length + rectangle.Breadth)
// }
//
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Length * rectangle.Breadth
// }
