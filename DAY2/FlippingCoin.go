package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	toss := rand.Intn(2)
	fmt.Println(toss)

	switch {
	case toss == 0:
		fmt.Println("TAIL")
	case toss == 1:
		fmt.Println("HEADS")
	}
}
