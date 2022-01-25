package string_sorting

import (
	"sort"
)

// CHALLENGE
// Sort a set of strings based on the following factors:
// 1- An odd length string should precede an even length string.
// 2- If both strings have odd lengths, the shorter of the two should precede.
// 3- If both strings have even lengths, the longer of the two should precede.
// 4- If the two strings have equal lengths, they should be in alphabetical order.
func customSorting(strArr []string) []string {
	result := make([]string, len(strArr))
	sort.Sort(customSort(strArr))
	copy(result, strArr)
	return result
}

type customSort []string

func (c customSort) Len() int {
	return len(c)
}

func (c customSort) Less(i, j int) bool {
	// same length -> alphabetical order
	if len(c[i]) == len(c[j]) {
		return c[i] < c[j]
	}
	first := len(c[i]) % 2
	second := len(c[j]) % 2
	// odd string before even string
	if first != second {
		return first > second
	}
	// both odd -> the shorter precedes
	if first == 1 {
		return len(c[i]) < len(c[j])
	}
	// both even -> the longer precedes
	return len(c[i]) > len(c[j])
}

func (c customSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
