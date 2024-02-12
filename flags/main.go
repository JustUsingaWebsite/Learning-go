package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count number of lines")
	flag.Parse()
	result := count(os.Stdin, *lines)
	fmt.Println(result)
}

func count(r io.Reader, linecount bool) int {
	scanner := bufio.NewScanner(r)

	if !linecount {
		scanner.Split(bufio.ScanWords)
	}

	counter := 0

	for scanner.Scan() {
		counter += 1
	}

	err := scanner.Err()
	if err != nil {
		return 0
	}
	return counter
}
