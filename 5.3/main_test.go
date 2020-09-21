package main

import (
	"testing"
	"time"
)

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := doWork(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected: %v, but received %v,", i, expected, r)
		}
		i++
	}
}

func TestDoWork_GeneratesAllNumbersB(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	const timeout = 2 * time.Second
	heartbeat, results := doWorkB(done, timeout/2, intSlice...)

	<-heartbeat

	i := 0
	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf(
					"index %v: expected: %v, but received %v,",
					i,
					expected,
					r,
				)
			}
			i++
		case <-heartbeat:
		case <-time.After(timeout):
			t.Fatal("test time out")
		}
	}
}
