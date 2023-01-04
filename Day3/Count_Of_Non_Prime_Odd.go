package main

import "fmt"

func Non_Prime_Odd(num int) {
	cnt := 0
	for i := num - 1; i > 1; i-- {
		ans := Is_Num_Prime(i)
		if ans == false {
			if i%2 != 0 {
				fmt.Printf("Non Prime Odd is: %d\n", i)
				cnt++
			}
		}
	}
	fmt.Printf("Count is :%d ", cnt)
}
func Is_Num_Prime(num int) bool {
	is_Prime := 0
	if num == 0 {
		//fmt.Printf("%d is Not Prime\n", num)
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				//fmt.Printf("%d is NOT Prime\n", num)
				is_Prime = 1
				return false
			}
		}
		if is_Prime == 0 {
			//fmt.Printf("%d is Prime\n", num)
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Println("Enter number: ")
	fmt.Scan(&n)
	Non_Prime_Odd(n)
}
