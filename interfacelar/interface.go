package main

import (
	"fmt"
	"math"
)

// why we need a interface in go?
// what is the use of interface in go?
// interface in go can be used to implement polymorphism, abstraction, and encapsulation.
// the use of interface in go is to define a set of methods that a type must implement.

// Shape is an interface that defines the area and perimeter methods
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle is a struct that represents a circle
type Circle struct {
	radius float64
}

// Rectangle is a struct that represents a rectangle
type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct{
	sideleft,sideright,base,height float64
}

func (c Triangle) area() float64{
	return (c.height*c.base)/2
}

func (c Triangle) perimeter()float64{
	return c.sideleft+c.sideright+c.base
}

// area returns the area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perimeter returns the perimeter of a circle
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// area returns the area of a rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// perimeter returns the perimeter of a rectangle
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// printShapeInfo prints the area and perimeter of a shape
func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func interfaces() {
	c := Circle{radius: 2.5}
	r := Rectangle{width: 3.0, height: 4.0}
	t:= Triangle{sideleft: 6,sideright: 6,height:4 ,base: 6}
	printShapeInfo(c)
	printShapeInfo(r)
	printShapeInfo(t)
}
