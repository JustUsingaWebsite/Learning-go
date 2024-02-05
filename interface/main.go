package main

import (
	"fmt"
	"math"
)

type myFraac struct {
	numer int
	denom int
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (m myFraac) String() string {
	return fmt.Sprintf("%d/%d", m.numer, m.denom)
}

type geometry interface {
	area() float64
	permin() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) permin() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) permin() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permin())
}

func main() {

	frac1 := myFraac{
		numer: 1,
		denom: 2,
	}

	fmt.Println(frac1)

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
