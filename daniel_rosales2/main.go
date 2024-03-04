package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func count(r io.Reader, bytesFlag bool, newFlag bool) (int, error) {
	data, err := ioutil.ReadAll(r) //read all bytes
	if err != nil {                // handle errors
		return -1, errors.New("error reading input")
	}
	if bytesFlag {
		//fmt.Println(data) <-- for testing ;)
		return len(data), nil
	}
	return 0, nil
}

func main() {

	byytes := flag.Bool("b", false, "number of bytes") // make bool flag
	newFlag := flag.Bool("n", false, "new flag")       // other bool flag
	flag.Parse()

	msg := os.Stdin // get user input (use CTRL + D if it keeps asking for info)

	bytesCount, err := count(msg, *byytes, *newFlag) // call count func

	if err != nil { // if there was an error
		fmt.Println("Error:", err)
	} else { // else show number of bytes
		fmt.Println("Number of bytes:", bytesCount)
	}
}
