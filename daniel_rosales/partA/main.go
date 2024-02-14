package main

import (
	"errors"
	"fmt"
	"math"
)

type circle struct {
	radius float64
} // made the easy struct

func (c *circle) area() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("radius cannot be negative")
	}
	return math.Pi * c.radius * c.radius, nil
} // simple check if it's negative if so return err

func (c *circle) circumference() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("radius cannot be negative")
	}
	return 2 * math.Pi * c.radius, nil
} // another simple check for negative numbers if true return err

func main() {

	//make a variable and call the area func
	//make sure to err check

	//make another variable and call circumference func
	//make sure to err check and handle it properly

	c := circle{radius: 5}

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
