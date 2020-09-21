package main

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	for val := range orDone(done, myChan) {

	}
}

// func main() {
// 	for val := range myChan {
//
// 	}
//
// loop:
// 	for {
// 		select {
// 		case <-done:
// 			break loop
// 		case maybeVal, ok := <-myChan:
// 			if ok == false {
// 				return
// 			}
// 		}
// 	}
// }
