package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	StateCapital := make(map[string]string, 10)
	StateCapital = map[string]string{
		"Maharashtra":    "Mumbai",
		"Karnataka":      "Bangaluru",
		"Goa":            "Panji",
		"Gujarat":        "Gandhinagar",
		"Tamil Nadu":     "Chennai",
		"Assam":          "Dispur",
		"Madhya Pradesh": "Bhopal",
		"Kerala":         "Thiruvananthapuram",
		"Rajasthan":      "Jaipur",
		"Bihar":          "Patna",
	}
	fmt.Printf("%v\n", StateCapital)
	j, err := json.MarshalIndent(StateCapital, " ", "\t")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}
}
