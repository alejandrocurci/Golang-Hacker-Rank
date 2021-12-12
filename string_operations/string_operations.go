package string_operations

import (
	"log"
	"regexp"
	"strings"
)

func ModifyString(str string) string {
	return reverse(removeDigits(strings.TrimSpace(str)))
}

func removeDigits(str string) string {
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "")
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
