package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Archived %v cycles of work before signalled to stop.\n", workCounter)
}

// func main() {
// 	start := time.Now()
// 	var c1, c2 <-chan int
// 	select {
// 	case <-c1:
// 	case <-c2:
// 	default:
// 		fmt.Printf("In default after %v\n\n", time.Since(start))
// 	}
// }

// func main() {
// 	var c <-chan int
// 	select {
// 	case <-c:
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Timed Out")
// 	}
// }

// func main() {
// 	c1 := make(chan interface{})
// 	close(c1)
// 	c2 := make(chan interface{})
// 	close(c2)
//
// 	var c1Count, c2Count int
// 	for i := 1000; i >= 0; i-- {
// 		select {
// 		case <-c1:
// 			c1Count++
// 		case <-c2:
// 			c2Count++
// 		}
// 	}
//
// 	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
// }

// func main() {
// 	start := time.Now()
// 	c := make(chan interface{})
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		close(c)
// 	}()
//
// 	fmt.Println("Blocking on read...")
// 	select {
// 	case <-c:
// 		fmt.Printf("Unblocked %v later.\n", time.Since(start))
// 	}
// }
