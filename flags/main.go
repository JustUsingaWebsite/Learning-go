package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	result := count(os.Stdin)
	fmt.Println(result)
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
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
