//1. print 3 favourite places in a line
//2. print 3 favourite places in different lines
//3. take user inputs for 3 favourite places

package main

import "fmt"

func main() {
	fmt.Println("Favourite Place1: Mumbai")
	fmt.Println("Favourite Place2: Bangalore")
	fmt.Print("Favourite Place3: Hyderabad")

	var Place1, Place2, Place3 string
	fmt.Printf("\nEnter 1st favourite place: ")
	fmt.Scanln(&Place1)
	fmt.Println("Enter 2nd favourite place: ")
	fmt.Scanln(&Place2)
	fmt.Println("Enter 3rd favourite place: ")
	fmt.Scanln(&Place3)
	println(Place1, Place2, Place3)

}
