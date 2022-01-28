package multiplier_generator

import "time"

const delay = 10 * time.Millisecond

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {
	num := -toAdd
	for {
		<-requestChan       // accept request for a new batch
		<-time.After(delay) // delay batch
		// generate batch and send it to channel
		for i := 0; i < batchSize; i++ {
			num += toAdd
			resultChan <- num
		}
	}
}
