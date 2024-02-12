package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello frm hello()")
}

func world() {
	fmt.Println("world frm world()")
}

func main() {
	go hello()
	go world()
	time.Sleep(time.Second)
}
