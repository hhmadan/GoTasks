package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-----UPPER-LOWER CASE PROGRAM-----")
	var place string
	fmt.Println("Enter most visited place by you")
	fmt.Scanln(&place)
	fmt.Println("String in UpperCase: ", strings.ToUpper(place))
	fmt.Println("String in LowerCase: ", strings.ToLower(place))

	fmt.Println("-----CONCATENATE PROGRAM-----")
	var strNum string
	fmt.Println("Enter Number: ")
	fmt.Scanln(&strNum)
	fmt.Println(strNum + "0") //+= is also valid

	fmt.Println("-----PRINT WITH SPACE IN BETWEEN-----")
	var fName, lName string
	fmt.Println("Enter First Name: ")
	fmt.Scanln(&fName)
	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&lName)

	fmt.Println(fName, lName)
}
