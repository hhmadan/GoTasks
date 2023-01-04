package main

import "fmt"

func Fibonacci_Series(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci_Series(n-1) + Fibonacci_Series(n-2)
}
func Fibonacci_Even_Series(n int) {
	_prev := 0
	_next := 1
	_cnt := 0

	if n <= 1 {
		fmt.Println(n)
	}
	for i := 2; i < n; i++ {
		_sum := _prev + _next
		if _sum%2 == 0 {
			//fmt.Println(" ", _sum)
			_cnt++
		}
		_prev = _next
		_next = _sum
	}
	fmt.Printf("Count of Even Numbers in Fibonacci Series: %d", _cnt)
}
func main() {
	var num int
	fmt.Println("Enter Number to get Fibonacci Series")
	fmt.Scan(&num)
	fmt.Println("Fibonacci Series is: ")
	for i := 0; i < num; i++ {
		fmt.Println(Fibonacci_Series(i))
	}
	//fmt.Println("Even Fibonacci Series is: ")
	Fibonacci_Even_Series(num)
}
