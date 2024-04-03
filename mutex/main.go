package main

import (
	"fmt"
	"sync"
)

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
