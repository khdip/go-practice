package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	Circle := circle{radius: 3}
	Rectangle := rectangle{width: 20, height: 12}

	fmt.Printf("Circle area: %f\n", getArea(Circle))
	fmt.Printf("Rectangle area: %f\n", getArea(Rectangle))
}
