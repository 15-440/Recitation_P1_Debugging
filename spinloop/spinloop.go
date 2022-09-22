// Adapted from codes used in 15440 Fall 2021 recitation

package main

import (
	"fmt"
	"log"
	"time"
)

func do() {
	// Nothing
}

func spinLoopConsumer(send chan struct{}, recv chan struct{}) {
	for {
		select {
		case <-send: // Get a request
			do()
		default: // No request, spinloop
			// Let's comment this (`default:``) out and run again!
			// Can you see difference?
		}
	}
}

func producer(send chan struct{}, recv chan struct{}) {
	start := time.Now()
	// Signal 100000 requst
	for i := 1; i < 100000; i++ {
		send <- struct{}{}
	}
	elapsed := time.Since(start)
	log.Printf("Time: %s", elapsed)
}

func main() {
	send := make(chan struct{})
	recv := make(chan struct{})
	go producer(send, recv)
	go spinLoopConsumer(send, recv)

	var input string
	fmt.Scan(&input)
}
