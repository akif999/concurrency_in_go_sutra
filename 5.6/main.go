package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	{
		type startGoroutineFn func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) (heartbeat <-chan interface{})

		newSteward := func(
			timeout time.Duration,
			startGroutine startGoroutineFn,
		) startGoroutineFn {
			return func(
				done <-chan interface{},
				pulseInterval time.Duration,
			) <-chan interface{} {
				heartbeat := make(chan interface{})
				go func() {
					defer close(heartbeat)

					var wardDone chan interface{}
					var wardHeartbeat <-chan interface{}
					startWard := func() {
						wardDone = make(chan interface{})
						wardHeartbeat = startGroutine(or(wardDone), timeout/2)
					}
					startWard()
					pulse := time.Tick(pulseInterval)

				monitorLoop:
					for {
						timeoutSignal := time.After(timeout)

						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case <-wardHeartbeat:
								continue monitorLoop
							case <-timeoutSignal:
								log.Println("stewward: ward unhealthy: restarting")
								close(wardDone)
								startWard()
								continue monitorLoop
							case <-done:
								return
							}
						}
					}
				}()

				return heartbeat
			}
		}

		log.SetOutput(os.Stdout)
		log.SetFlags(log.Ltime | log.LUTC)

		doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
			log.Println("ward: Hello, I'm irresponsible!")
			go func() {
				<-done
				log.Println("ward: I am halting.")
			}()
			return nil
		}
		doWorkWithSteward := newSteward(4*time.Second, doWork)

		done := make(chan interface{})
		time.AfterFunc(9*time.Second, func() {
			log.Println("main: halting stewward and ward.")
			close(done)
		})

		for range doWorkWithSteward(done, 4*time.Second) {
		}
		log.Println("Done")
	}

	doWorkFn := func(
		done <-chan interface{},
		intList ...int,
	) (startGoroutineFn, <-chan interface{}) {
		intChanStream := make(chan (<-chan interface{}))
		intStream := bridge(done, intChanStream)
		doWork := func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			intStream := make(chan interface{})
			heartbeat := make(chan interface{})
			go func() {
				defer close(intStream)
				select {
				case intChanStream <- intStream:
				case <-done:
					return
				}

				pulse := time.Tick(pulseInterval)

				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Printf("negative value: %v\n", intVal)
							return
						}

						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}
						}
					}
				}
			}()
			return heartbeat
		}
		return doWork, intStream
	}

	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)

	for intVal := range take(done, intStream, 6) {
		fmt.Printf("Received: %v\n", intVal)
	}
}
