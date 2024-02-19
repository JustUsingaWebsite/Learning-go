package main

import (
	"flag"
	"fmt"
)

func main() {
	msg := flag.String("text", "Hello there!", "messsage to display")
	flag.Parse()
	fmt.Println(*msg)
}
