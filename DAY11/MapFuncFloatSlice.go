package main

import "fmt"

func imap(intSlice []float64, f func(float64) float64) []float64 {
	mapped := make([]float64, len(intSlice))

	for i, num := range intSlice {
		mapped[i] = f(num)
	}
	return mapped
}
func main() {
	intSlice := []float64{0.6, 0.3, 0.84, 0.04}

	percentSlice := imap(intSlice, func(float float64) float64 {
		return float * 100
	})

	halfValueSlice := imap(intSlice, func(float float64) float64 {
		return float / 2
	})

	fmt.Println(percentSlice)
	fmt.Println(halfValueSlice)
}
