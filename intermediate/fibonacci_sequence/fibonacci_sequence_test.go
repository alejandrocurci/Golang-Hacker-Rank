package fibonacci_sequence

import (
	"fmt"
	"reflect"
	"testing"
)

type test struct {
	skipped  int
	amount   int
	expected []int
}

func TestModuloFibonacciSequence(t *testing.T) {
	cases := []test{
		{
			skipped:  0,
			amount:   6,
			expected: []int{1, 2, 3, 5, 8, 13},
		},
		{
			skipped:  50,
			amount:   4,
			expected: []int{951280099, 316291173, 267571272, 583862445},
		},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprintf("Test for %v skipped, %v results", tt.skipped, tt.amount), func(t *testing.T) {
			boolCh := make(chan bool)
			resultCh := make(chan int)
			actualResults := make([]int, 0, tt.amount)
			go ModuloFibonacciSequence(boolCh, resultCh)
			for i := 0; i < tt.amount+tt.skipped; i++ {
				boolCh <- true
				num := <-resultCh
				if i >= tt.skipped {
					actualResults = append(actualResults, num)
				}
			}
			if !reflect.DeepEqual(tt.expected, actualResults) {
				t.Errorf("expected %v, got %v", tt.expected, actualResults)
			}
		})
	}
}

// check no negative numbers are coming as results
func TestNegativeResults(t *testing.T) {
	cases := []test{
		{
			skipped: 100,
			amount:  200,
		},
	}

	for _, tt := range cases {
		t.Run("zero negative numbers", func(t *testing.T) {
			boolCh := make(chan bool)
			resultCh := make(chan int)
			negativeNumbers := make([]int, 0)
			go ModuloFibonacciSequence(boolCh, resultCh)
			for i := 0; i < tt.amount+tt.skipped; i++ {
				boolCh <- true
				num := <-resultCh
				if i >= tt.skipped {
					if num < 0 {
						negativeNumbers = append(negativeNumbers, num)
					}
				}
			}
			if len(negativeNumbers) > 0 {
				t.Errorf("got %v negative numbers", len(negativeNumbers))
			}
		})
	}
}
