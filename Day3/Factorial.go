package main

import "fmt"

func Recursive_Factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Recursive_Factorial(num-1)
}
func Iterative_Factorial(num int) int {
	fact := 1
	for num >= 1 {
		fact *= num
		num -= 1
	}
	return fact
}
func main() {
	var input_Num int
	fmt.Println("Enter Number to get its Factorial: ")
	fmt.Scan(&input_Num)
	r_fac := Recursive_Factorial(input_Num)
	i_fac := Iterative_Factorial(input_Num)

	fmt.Printf("Factorial of %d using Recursion is %d\n", input_Num, r_fac)
	fmt.Printf("Factorial of %d using Iteration is %d", input_Num, i_fac)
}
