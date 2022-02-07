package fibonacci_sequence

import (
	"time"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	mod := 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10
	a, b := 0, 1
	for range requestChan {
		a, b = b, (a+b)%mod
		select {
		case <-time.After(10 * time.Millisecond):
			resultChan <- b
		}
	}
}
