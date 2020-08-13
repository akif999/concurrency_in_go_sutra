package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Recieved: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

// func main() {
// 	var stdoutBuff bytes.Buffer
// 	defer stdoutBuff.WriteTo(os.Stdout)
//
// 	intStream := make(chan int, 4)
// 	go func() {
// 		defer close(intStream)
// 		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
// 		for i := 0; i < 5; i++ {
// 			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
// 			intStream <- i
// 		}
// 	}()
//
// 	for integer := range intStream {
// 		fmt.Fprintf(&stdoutBuff, "Recieved %v.\n", integer)
// 	}
// }

// func main() {
// 	begin := make(chan interface{})
// 	var wg sync.WaitGroup
// 	for i := 0; i <= 4; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			<-begin
// 			fmt.Printf("%v has begun\n", i)
// 		}(i)
// 	}
//
// 	fmt.Println("Unblocking goroutines...")
// 	close(begin)
// 	wg.Wait()
// }

// func main() {
// 	intStream := make(chan int)
// 	go func() {
// 		defer close(intStream)
// 		for i := 1; i <= 5; i++ {
// 			intStream <- i
// 		}
// 	}()
//
// 	for integer := range intStream {
// 		fmt.Printf("%v", integer)
// 	}
// }

// func main() {
// 	intStream := make(chan int)
// 	close(intStream)
// 	integer, ok := <-intStream
// 	fmt.Printf("(%v): %v", ok, integer)
// }

// func main() {
// 	stringStream := make(chan string)
// 	go func() {
// 		stringStream <- "Hello channels!"
// 	}()
// 	salutation, ok := <-stringStream
// 	fmt.Printf("(%v): %v", ok, salutation)
// }

// func main() {
// 	stringStream := make(chan string)
// 	go func() {
// 		if 0 != 1 {
// 			return
// 		}
// 		stringStream <- "Hello channels!"
// 	}()
// 	fmt.Println(<-stringStream)
// }

//func main() {
//	stringStream := make(chan string)
//	go func() {
//		stringStream <- "Hello channels!"
//	}()
//	fmt.Println(<-stringStream)
//}
