package fibonacci_sequence

import (
	"math"
	"time"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	a, b := 0, 1
	for range requestChan {
		a, b = b, a+b
		select {
		case <-time.After(10 * time.Millisecond):
			mod := int(math.Pow10(9))
			resultChan <- b % mod
		}
	}
}
