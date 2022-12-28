package main

import "fmt"

func main() {
	harmonic := 1.00
	var n int
	fmt.Println("Enter Number: ")
	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		harmonic += (1 / (float64)(i))
	}
	fmt.Println(harmonic)
}
