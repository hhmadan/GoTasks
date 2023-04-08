package main

import (
	"fmt"
)

func main() {
	state_Capital := make(map[string]string)

	fmt.Printf("Data Type of state_Capital is: %T\n\n", state_Capital)
	state_Capital["Maharashtra"] = "Mumbai"
	state_Capital["Karnataka"] = "Bangalore"
	state_Capital["Goa"] = "Panji"
	state_Capital[" Telengana"] = "Hyderabad"
	state_Capital["Tamil nadu"] = "Chennai"
	state_Capital["Gujrat"] = "Ga ndhinagar"
	state_Capital["Kerala"] = "Thiruvananthapuram"
	state_Capital["Rajasthan"] = "Jaipur"
	state_Capital["Assam"] = "Dispur"

	// iterating over map
	for state, capital := range state_Capital {
		fmt.Printf("%s ==> %s\n", state, capital)
	}
}
