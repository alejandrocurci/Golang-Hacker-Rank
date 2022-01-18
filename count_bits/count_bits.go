package count_bits

import (
	"strconv"
	"strings"
)

// CHALLENGE
// Implement a function that counts the number of set bits in the binary representation
// of a 32-bit integer
// Examples: 126 has 6 bits set (1111110 in binary)
// 128 has 1 bit set (10000000 in binary)

func countBits(num uint32) int32 {
	b := convertToBinary(num)
	return int32(strings.Count(b, "1"))
}

func convertToBinary(num uint32) string {
	var binary string
	for num > 0 {
		binary = strconv.Itoa(int(num%2)) + binary
		num = (num - num%2) / 2
	}
	return binary
}
