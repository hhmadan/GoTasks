package main

import "fmt"

func multiplesSum_Of_Three_n_Five(num int) {
	sum := 0
	for i := num - 1; i >= 1; i-- {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			fmt.Print(i, " ")
		}
	}
	fmt.Printf("\nSum of all Multiples is: %d", sum)
}
func main() {
	var num int
	fmt.Print("Enter Number: ")
	fmt.Scan(&num)
	multiplesSum_Of_Three_n_Five(num)
}
