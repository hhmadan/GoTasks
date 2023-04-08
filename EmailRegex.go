package main

import (
	"fmt"
	"regexp"
)

func main() {
	var email string
	fmt.Println("Enter Email-Id: ")
	fmt.Scanln(&email)

	pattern := "[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
	match, _ := regexp.MatchString(pattern, email)
	fmt.Println("Match: ", match)
}
