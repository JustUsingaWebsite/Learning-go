package main

import (
	"fmt"
)

type animal struct {
	alive     bool
	breathing bool
	walking   bool
	name      string
}

type rectangle struct {
	length float64
	WIDTH  float64
}

func (a animal) bark() string {
	return a.name + " has barked"
}

func showDAtaof(value animal) {
	fmt.Printf("%T", value)
	return
}

func (r rectangle) area() float64 {
	return r.length * r.WIDTH
}

func main() {

	var child = [3]string{
		"bingo", "cat", "dog",
	}

	for true {
		fmt.Println(child[0])
		break
	}

	var jag animal
	jag.walking = false
	jag.name = "jaguar"
	jag.alive = true
	fmt.Println(jag.bark())
	showDAtaof(jag)

	cat := animal{
		alive:     false,
		breathing: false,
		walking:   false,
		name:      "cat",
	}

	showDAtaof(cat)

	rec1 := rectangle{
		WIDTH:  4,
		length: 46,
	}

	fmt.Println("\n"+"----------------------------------------", "\n")
	fmt.Println(rec1.area())
}
