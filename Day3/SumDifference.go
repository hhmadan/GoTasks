package main

import "fmt"

func sum(n1, n2 int) {
	fmt.Printf("Sum is: %d + %d = %d\n", n1, n2, (n1 + n2))
}
func diff(n1, n2 int) {
	if n1 > n2 {
		fmt.Printf("Difference is: %d - %d = %d\n", n1, n2, (n1 - n2))
	} else {
		diff := n2 - n1
		fmt.Printf("Difference is: | %d - %d | = %d\n", n1, n2, diff)
	}
}
func sum_Diff(n1, n2 int) {
	fmt.Printf("Sum is: %d + %d = %d\n", n1, n2, (n1 + n2))
	if n1 > n2 {
		fmt.Printf("Difference is: %d - %d = %d\n", n1, n2, (n1 - n2))
	} else {
		diff := n2 - n1
		fmt.Printf("Difference is: | %d - %d | = %d\n", n1, n2, diff)
	}
}
func main() {
	sum(10, 10)
	diff(5, 2)
	diff(4, 5)
	sum_Diff(10, 6)
}
