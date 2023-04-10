package main

import "fmt"

func sum(num ...int) float64 {
	total, cnt := 0, 0
	for _, num := range num {
		total += num
		cnt++
	}
	return float64(total) / float64(cnt)
}
func main() {
	s1 := sum(2, 5, 1, 4)
	s2 := sum(2, 5, 1, 4, 8)
	s3 := sum(2, 5, 1, 4, 2, 1)
	fmt.Printf("%.2f %.2f %.2f", s1, s2, s3)
}
