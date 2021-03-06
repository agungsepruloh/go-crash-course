package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
}

// Circle interface
type Circle struct {
	x, y, radius float64
}

// Rectangle interface
type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getAarea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f \n", getAarea(circle))
	fmt.Printf("Rectangle area: %f \n", getAarea(rectangle))
}
