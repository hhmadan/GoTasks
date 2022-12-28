package main

import "fmt"

func main() {
	var choice int
	var tc, tf float32
	fmt.Printf("\n**WELCOME TO TEMPERATURE UNIT CONVERTER**\n")
	fmt.Printf("\n----CONVERSION MENU----\n")
	fmt.Printf("1.Centigrade to Farenheit\n")
	fmt.Printf("2.Farenheit to Centigrade\n")
	fmt.Printf("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("\nEnter Temperature in Centigrade:")
		fmt.Scan(&tc)
		tf = (tc * 1.8) + 32
		fmt.Printf("Temperature in Farenheit: %f F", tf)
	case 2:
		fmt.Printf("\nEnter Temperature in Farenheit:")
		fmt.Scan(&tf)
		tc = (tf - 32) * 0.55
		fmt.Printf("Temperature in Centrigrade: %f C", tc)
	default:
		fmt.Println("INVALID CHOICE..!")
	}
}
