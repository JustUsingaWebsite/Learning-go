package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var cchar [2]int

	for i := 0; i < len(cchar); i++ {
		cchar[i] = rand.Intn(100)
	}

	for true {
		fmt.Println("finished")
		break
	}
	//type conversion

	x := 12.3
	y := 7
	z := x + float64(y)

	fmt.Println(cchar, z)
}
