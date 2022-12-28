package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	interest := 7.50
	t := 2.5

	final_amt := principal * (1 + interest/t) * interest * t

	fmt.Println(final_amt)

	var P, Y, R float64
	//P= Principal Amount
	//Y= Number of Years
	//R= Rate of Interest
	var r, n float64
	P = 1000.0
	Y = 3.5
	R = 7.5
	r = (R) / (12 * 100)
	n = 12 * (Y)
	pay := (P * r) / (1 - (math.Pow((1 + r), -n)))

	fmt.Printf("%f", pay)
}
