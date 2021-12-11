package count_bits

import (
	"strconv"
	"strings"
)

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
