package main

import (
	"fmt"
	"regexp"
)

func main() {
	var password string
	fmt.Println("Enter Password: ")
	fmt.Scanln(&password)

	pattern := `(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d^a-zA-Z0-9].{5,50}$`
	match, _ := regexp.MatchString(pattern, password)
	fmt.Println("Match: ", match)
}
