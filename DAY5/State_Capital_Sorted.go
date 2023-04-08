package main

import (
	"fmt"
	"sort"
)

func main() {
	state_Capital := make(map[string]string)
	fmt.Printf("Data Type of state_Capital is: %T\n\n", state_Capital)
	state_Capital["Maharashtra"] = "Mumbai"
	state_Capital["Karnataka"] = "Bangalore"
	state_Capital["Goa"] = "Panji"
	state_Capital[" Telengana"] = "Hyderabad"
	state_Capital["Tamil nadu"] = "Chennai"
	state_Capital["Gujrat"] = "Gandhinagar"
	state_Capital["Kerala"] = "Thiruvananthapuram"
	state_Capital["Rajasthan"] = "Jaipur"
	state_Capital["Assam"] = "Dispur"

	for state, capital := range state_Capital {
		fmt.Printf("%s ==> %s\n", state, capital)
	}
	var _state []string

	for s := range state_Capital {
		_state = append(_state, s)
	}
	sort.Strings(_state)
	fmt.Println(_state)
}
