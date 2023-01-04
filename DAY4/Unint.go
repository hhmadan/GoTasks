package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := []int{}
	var input string
	var size int
	var numbers []int

	fmt.Println("Enter size")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		fmt.Println("Please enter a number")
		fmt.Scan(&input)
		if input == "X" {
			break
		}

		sliceVar, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input")
			continue
		}
		slice = append(slice, sliceVar)
		sort.Ints(slice)
		fmt.Println(slice)
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			numbers = append(numbers, slice[i])
		}

	}
	fmt.Println(numbers)
}
