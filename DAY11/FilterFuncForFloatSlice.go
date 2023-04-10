package main

import (
	"fmt"
	"math"
)

func filter(intSlice []float64, getZeroDecimal func(float64) bool) []float64 {
	var zerodecimalSlice []float64
	for _, val := range intSlice {
		if getZeroDecimal(val) {
			zerodecimalSlice = append(zerodecimalSlice, val)
		}

	}
	return zerodecimalSlice
}
func main() {
	floatSlice := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}

	zerodecimalSlice := filter(floatSlice, func(f float64) bool {
		return math.Mod(f, 1) == 0
	})

	factorOfThree := filter(floatSlice, func(f float64) bool {
		return math.Mod(f, 3.0) == 0
	})
	fmt.Println(zerodecimalSlice)
	fmt.Println(factorOfThree)
}
