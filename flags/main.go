package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	msg := flag.String("text", "Hello there!", "messsage to display")

	caps := flag.Bool("c", false, "should text be uppercase")
	flag.Parse()

	if *caps {
		*msg = strings.ToUpper(*msg)
	}

	fmt.Println(*msg)
}
