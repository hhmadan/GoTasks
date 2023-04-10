package main

import (
	"fmt"
	"strings"
)

func filterFunc(inputString []string, f func(string) bool) []string {
	prefixSlice := []string{} //temp. slice only fname
	for _, words := range inputString {
		if f(words) { //if anony. func returns true...
			prefixSlice = append(prefixSlice, words)
		}
	}
	return prefixSlice
}
func main() {
	inputString := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	checkPrefix := filterFunc(inputString, func(s string) bool {
		return strings.HasPrefix(s, "b")
		//return s[0] == 'b'
	})

	threeLetters := filterFunc(inputString, func(s string) bool {
		return len(s) == 3
	})
	fmt.Println(checkPrefix)
	fmt.Println(threeLetters)
}
