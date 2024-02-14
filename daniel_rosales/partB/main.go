package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
} // triangle struct

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
} // make area method for triangle and return half base times height

func (t triangle) perimeter() float64 {
	return t.base + t.height
} // make perimeter method for triangle and return calculations

type geometry interface {
	area() float64
	perimeter() float64
} // make type interface to make triangle a type of geometry

func measure(g geometry) {
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
} // measure func to make traingle a type of geometry with area n perimeter

func main() {

	t := triangle{base: 3, height: 4}
	//decalere and run func to check that triangle is a type of geometry
	measure(t)
}
