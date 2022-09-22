// Adapted from codes used in 15440 Fall 2021 recitation

package main

func consumer(ch chan int) {
	for {
		select {
		case <-ch:
			// do nothing
		}
	}
}

func do() {
	for {
		// do nothing
	}
}

func main() {
	ch := make(chan int)
	//go do() // go only produce error message when all routines are dead, try to uncomment this
	ch <- 1
	go consumer(ch)
	// fmt.Println(<-ch)
}
