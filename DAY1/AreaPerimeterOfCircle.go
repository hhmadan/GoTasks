package main

import "fmt"

func main() {
	var area, r float32
	fmt.Println("Enter radius")
	fmt.Scanln(&r)
	area = 3.14 * r * r
	circum := 2 * 3.14 * r
	fmt.Println("Area of Circle: ", area)
	fmt.Println("Perimeter of Circle: ", circum)
}
