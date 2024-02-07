package main

import (
	"errors"
	"fmt"
	"math"
)

type triange struct {
	base   float64
	height float64
}

type cricle struct {
	radius float64
}

type Shapes interface {
	area() float64
	perimeter() float64
}

func (t triange) area() float64 {
	return t.base * t.height
}

func (t triange) perimeter() float64 {
	return 2 * (t.base + t.height)
}

func (c cricle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func square(s float32) (float64, float64, error) {

	if s <= 0 {
		err := errors.New("The sies of the square cannot be 0 or below")
		return -1, -1, err
	} else {
		cc1 := triange{3, 3}
		return cc1.area(), cc1.perimeter(), nil
	}
}

func testTriangle(val triange) {

	a := val.area()
	p := val.perimeter()

	if a < 0 || p < 0 {
		fmt.Println("Answer cannot be under zero: error")
		return
	}

	fmt.Println("area of the triangle:", val.area())
	fmt.Println("perimeter of the triangle:", val.perimeter())
	fmt.Println("----------------------------")
}

func testCircle(val cricle) {

	r := val.area()

	if r < 0 {
		fmt.Println("Answer cannot be under zero: error")
		return
	}

	fmt.Println("area of the circle:", val.area())
	fmt.Println("----------------------------")
}

func testsquare(r1 float64, r2 float64, err error) {

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Area and perimeter:", r1, r2)
}

func main() {
	t1 := triange{5, 4}
	c1 := cricle{15}
	result1, result2, err := square(4)

	testTriangle(t1)
	testCircle(c1)
	testsquare(result1, result2, err)
}
