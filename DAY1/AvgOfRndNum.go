package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Generating 5 random numbers...")
	rand.Seed(time.Now().UnixNano())

	var ran [5]int
	sum := 0
	min := 10
	max := 50

	for i := 0; i < 5; i++ {
		ran[i] = rand.Intn((max - min) + min)
		sum += ran[i]
	}
	fmt.Println(ran)
	fmt.Println("Sum of Random numbers is: ", sum)
	fmt.Println("Average of Random Numbers is: ", (sum / (len(ran))))
}
