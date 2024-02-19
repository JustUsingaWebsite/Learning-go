package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	msg := flag.String("text", "Hello there!", "messsage to display")

	caps := flag.Bool("c", false, "should text be uppercase")
	nums := flag.Int("n", 1, "number of times is displayed")
	flag.Parse()

	if *caps {
		*msg = strings.ToUpper(*msg)
	}

	for i := 0; i < *nums; i++ {
		fmt.Println(*msg)
	}

}
