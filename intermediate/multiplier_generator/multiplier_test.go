package multiplier_generator

import (
	"fmt"
	"reflect"
	"testing"
)

type test struct {
	batchSkipped int
	batchPrint   int
	batchSize    int
	toAdd        int
	expected     [][]int
}

func TestBurstyRateLimiter(t *testing.T) {
	cases := []test{
		{
			batchSkipped: 0,
			batchPrint:   4,
			batchSize:    2,
			toAdd:        3,
			expected:     [][]int{{0, 3}, {6, 9}, {12, 15}, {18, 21}},
		},
		{
			batchSkipped: 1,
			batchPrint:   3,
			batchSize:    3,
			toAdd:        2,
			expected:     [][]int{{6, 8, 10}, {12, 14, 16}, {18, 20, 22}},
		},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprintf("Test for batchSkipped=%v, batchPrint=%v, batchSize=%v, toAdd=%v",
			tt.batchSkipped, tt.batchPrint, tt.batchSize, tt.toAdd), func(t *testing.T) {
			boolCh := make(chan bool)
			resultCh := make(chan int)
			actualResults := make([][]int, 0, tt.batchPrint)
			go BurstyRateLimiter(boolCh, resultCh, tt.batchSize, tt.toAdd)
			for i := 0; i < tt.batchPrint+tt.batchSkipped; i++ {
				boolCh <- true
				batch := make([]int, 0)
				for j := 0; j < tt.batchSize; j++ {
					num := <-resultCh
					batch = append(batch, num)
				}
				if i >= tt.batchSkipped {
					actualResults = append(actualResults, batch)
				}
			}
			if !reflect.DeepEqual(tt.expected, actualResults) {
				t.Errorf("expected %v, got %v", tt.expected, actualResults)
			}
		})
	}
}
