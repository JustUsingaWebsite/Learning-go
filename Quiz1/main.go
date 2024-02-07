package main

import (
	"fmt"
	"math"
)

func square(x float64) float64 {
	return math.Pow(x, 2)
}

type pet struct {
	name  string
	breed string
	age   int
}

func main() {
	result := square(5.0)
	fmt.Println("Squared: ", result)

	pet1 := pet{name: "bowser", breed: "labrador", age: 5}
	pet2 := pet{name: "tan tan", breed: "persian", age: 6}

	fmt.Printf("Pet1\n %v\n", pet1)
	fmt.Printf("Pet2\n %v", pet2)
}
