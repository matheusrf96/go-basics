package main

import (
	"fmt"
	"math"
)

// Interface shape
type shape interface {
	area() float64
}

func writeArea(s shape) {
	fmt.Printf("The area of the form is %0.2f\n", s.area())
}

// Rectangle
type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

// Circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangle{5, 3}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
