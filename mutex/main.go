package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mx := sync.Mutex{}
	mx.Lock()

	go func() {
		mx.Lock()
		fmt.Println("hola amigo")
		mx.Unlock()
	}()

	fmt.Println("hello friend")
	mx.Unlock()
	time.Sleep(time.Second)
}
