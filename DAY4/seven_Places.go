package main

import "fmt"

func main() {

	var places_List [7]string

	for place := range places_List {
		fmt.Println("Enter your favouriate place ")
		fmt.Scanln(&places_List[place])
	}
	fmt.Println(places_List)
	for word := 0; word < len(places_List); word++ {
		letterCount := 0
		for letter := 0; letter < len(places_List[word]); letter++ {
			letterCount++
		}
		fmt.Println("Count of word in entered place ", places_List[word], "is : ", letterCount)
	}
}
