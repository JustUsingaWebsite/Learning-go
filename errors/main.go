package main

import (
	"errors"
	"fmt"
	"reflect"
)

var myvar interface{} = 12.34;
	if _, ok := myvar.(float64); ok{
		fmt.Println("it's a float")
	}

func area(length, width float64) (float64, error) {
	if length < 0 {
		err := errors.New("length or width must be 0 or above")
		return -1.0, err
	}

	if width < 0 {
		err := errors.New("width must be 0 or above")
		return -1.0, err
	}

	return length * width, nil
}

func main() {
	result, err := area(15, -2.5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	
}
