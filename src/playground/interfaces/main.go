package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

type AngledShape struct {
	base   float64
	height float64
}

func (a AngledShape) Area() float64 {
	return a.base * a.height
}

func (a AngledShape) Perimeter() float64 {
	return (a.base + a.height) * 2
}

func (a AngledShape) String() string {
	return fmt.Sprintf("Shape area is: %.2f and perimeter is: %.2f", a.Area(), a.Perimeter())
}

type Triangle struct {
	AngledShape
}
type Rectangle struct {
	AngledShape
}

const PI = 3.141592653589793

type Circle struct {
	radius float64
}

func (s Circle) Area() float64 {
	return s.radius * s.radius * PI
}

func (s Circle) Perimeter() float64 {
	return 2 * (s.radius * PI)
}

func (s Circle) String() string {
	return fmt.Sprintf("Shiny cicle in the sky, its area is : %.2f its perimeter is %.2f", s.Area(), s.Perimeter())
}

func ShapeCalculator(s []Shape) (float64, float64) {
	var totalArea float64 = 0
	var totalPerimenter float64 = 0
	for _, ss := range s {
		totalArea += ss.Area()
		totalPerimenter += ss.Perimeter()
	}
	return totalArea, totalPerimenter
}

func main() {

	t := Triangle{AngledShape{base: 33, height: 44}}
	r := Rectangle{AngledShape{base: 44, height: 55}}
	c := Circle{radius: 15}

	shapes := []Shape{t, r, c}
	totalArea, totalPerimeter := ShapeCalculator(shapes)
	fmt.Printf("total shape are is %.2f and total perimeter is %.2f", totalArea, totalPerimeter)

	for _, s := range shapes {
		fmt.Println(s)
	}
}
