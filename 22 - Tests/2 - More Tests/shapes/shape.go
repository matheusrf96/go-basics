package shapes

import (
	"math"
)

// Interface shape
type Shape interface {
	Area() float64
}

// Rectangle
type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// Circle
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
