package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var factors_of_num []string
	fmt.Println("Enter Number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			str := strconv.Itoa(i)
			factors_of_num = append(factors_of_num, str)
		}
	}
	//fmt.Println(factors_of_num)
	for i := 0; i < len(factors_of_num); i++ {
		if factors_of_num[i] == "3" {
		} else if factors_of_num[i] == "3" {
			fmt.Println("Pling")
		} else if factors_of_num[i] == "5" {
			fmt.Println("Plang")
		} else if factors_of_num[i] == "7" {
			fmt.Println("Plong")
		}
	}
}
