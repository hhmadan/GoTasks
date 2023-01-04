package main

import "fmt"

func main() {
	var arr = [5][2]int{}
	arr[0][0] = 10
	arr[1][0] = 12
	arr[2][0] = 15
	arr[3][0] = 19
	arr[4][0] = 24

	for i := 0; i < len(arr); i++ {
		arr[i][1] = arr[i][0] * 2
	}
	fmt.Println(arr)
}
