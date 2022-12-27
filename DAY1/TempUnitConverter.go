package main

import "fmt"

func main() {
	var tc, tf float32
	fmt.Println("Enter Temperature in Centigrade")
	fmt.Scanln(&tc)

	tf = (tc * 1.8) + 32
	fmt.Println("Temperature in Centrigrade: ", tc)
	fmt.Println("Temperature in Farenheit: ", tf)
}
