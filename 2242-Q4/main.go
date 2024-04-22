package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// creating a WaitGroup to wait for both goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// greating two channels, ch1 for Player 1 and ch2 for Player 2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine for Player 1
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+600) * time.Microsecond)
		// sending a message to ch1
		ch1 <- "Player's 1 buzzer."
		// signaling WaitGroup that this goroutine has finished
		wg.Done()
	}()

	// goroutine for Player 2
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Microsecond)
		// sending a message to ch2
		ch2 <- "Player's 2 buzzer."
		// signaling WaitGroup that this goroutine has finished
		wg.Done()
	}()

	// goroutine to wait for both players to finish and then close the channels
	go func() {
		wg.Wait() // wait for both players
		close(ch1)
		close(ch2)
	}()

	// loop to select messages from ch1 and ch2
	for {
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println(msg1) // print player 1's message if received
			}
		case msg2, ok := <-ch2:
			if ok {
				fmt.Println(msg2) // print player 2's message if received
			}
		}
		// break the loop when both channels are closed
		if len(ch1) == 0 && len(ch2) == 0 {
			break
		}
	}
}
