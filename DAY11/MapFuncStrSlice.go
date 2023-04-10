package main

import (
	"fmt"
	"strings"
)

func imap(data []string, f func(string) string) []string {

	mapped := make([]string, len(data))

	for i, word := range data {

		mapped[i] = f(word)
		fmt.Println(i, " ", word, " ", mapped)
	}
	return mapped
}

func main() {
	strSlice := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	initCapital := imap(strSlice, func(i string) string {
		return strings.Title(i)
	})

	addColon := imap(strSlice, func(i string) string {
		return i + ":"
	})

	fmt.Println(initCapital)
	fmt.Println(addColon)
}
