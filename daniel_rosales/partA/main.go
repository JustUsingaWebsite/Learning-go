package main

import (
	"errors"
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c *circle) area() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("radius cannot be negative")
	}
	return math.Pi * c.radius * c.radius, nil
}

func (c *circle) circumference() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("radius cannot be negative")
	}
	return 2 * math.Pi * c.radius, nil
}

func main() {

	c := circle{radius: 5.0}

	area, err := c.area()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Area of the circle:", area)
	}

	circumference, err := c.circumference()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Circumference of the circle:", circumference)
	}
}
