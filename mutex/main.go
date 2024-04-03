package main

import (
	"fmt"
	"sync"
)

func deposit(bal *int, amt int, mx *sync.Mutex, wg *sync.WaitGroup) {
	mx.Lock()
	*bal = *bal + amt
	mx.Unlock()
	wg.Done()
}

func withdraw(bal *int, amt int, mx *sync.Mutex, wg *sync.WaitGroup) {
	mx.Lock()
	*bal = *bal - amt
	mx.Unlock()
	wg.Done()
}

func main() {
	var mx sync.Mutex
	var wg sync.WaitGroup

	bal := 1000

	wg.Add(2)

	go deposit(&bal, 500, &mx, &wg)
	withdraw(&bal, 1000, &mx, &wg)

	wg.Wait()
	fmt.Println(bal)
}
