package remainder_sorting

import (
	"sort"
	"unicode/utf8"
)

// Challenge
// Implement a function that receives an array of strings and sort them based on the following heuristics:
// 1- the primary sort is by increasing remainder of the strings length, modulo 3
// 2- if multiple strings have the same remainder, they should be sorted in alphabetical order
// E.g: strArr=["Colorado","Utah","Wisconsin","Oregon"] should return sortedArray =["Oregon","Wisconsin","Utah","Colorado"]

func RemainderSorting(list []string) []string {
	result := make([]string, len(list))
	sort.Sort(remainderSort(list))
	copy(result, list)
	return result
}

// remainderSort implements sort.Interface
type remainderSort []string

func (r remainderSort) Len() int {
	return len(r)
}

func (r remainderSort) Less(i, j int) bool {
	if len(r[i])%3 == len(r[j])%3 {
		iRune, _ := utf8.DecodeRuneInString(r[i])
		jRune, _ := utf8.DecodeRuneInString(r[j])
		return int32(iRune) < int32(jRune)
	}
	return len(r[i])%3 < len(r[j])%3
}

func (r remainderSort) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
