package main

import "fmt"

func main() {
	var sizeOfArr int
	fmt.Println("Enter Size of Array: ")
	fmt.Scan(&sizeOfArr)
	minMaxArray := make([]int, sizeOfArr)

	fmt.Println("----Input the Elements in Array----")
	for element := range minMaxArray {
		fmt.Scan(&minMaxArray[element])
	}
	fmt.Println(minMaxArray)
	for index := 0; index < sizeOfArr; index++ {
		for index1 := index + 1; index1 < sizeOfArr; index1++ {
			if minMaxArray[index] > minMaxArray[index1] {
				temp := minMaxArray[index]
				minMaxArray[index] = minMaxArray[index1]
				minMaxArray[index1] = temp
			}
		}
	}
	fmt.Println("Largest Number is: ", minMaxArray[sizeOfArr-1])
	fmt.Println("Second Largest Number is: ", minMaxArray[sizeOfArr-2])
	fmt.Println("Smallest Number is: ", minMaxArray[0])
	fmt.Println("Second Smallest Number is: ", minMaxArray[1])
}
