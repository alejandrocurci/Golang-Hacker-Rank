package count_bits

import (
	"fmt"
	"testing"
)

type cases struct {
	input  uint32
	binary string
	bits   int32
}

func TestCases(t *testing.T) {
	tests := []cases{
		{3, "11", 2},
		{10, "1010", 2},
		{126, "1111110", 6},
		{127, "1111111", 7},
		{128, "10000000", 1},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("Number %v", tt.input)
		t.Run(name, func(t *testing.T) {
			b := convertToBinary(tt.input)
			c := countBits(tt.input)
			if b != tt.binary {
				t.Errorf("binary expected %s, actual %s", tt.binary, b)
			}
			if c != tt.bits {
				t.Errorf("count expected %v, actual %v", tt.bits, c)
			}
		})
	}
}
