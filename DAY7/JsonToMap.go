package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonAppleReceipe := `{
		"Thinly SLiced Peeled Apples":"6 Cups",
		"sugar":"3/4 cup",
		"flour":"2 tablespoons",
		"cinamon":"1/4 teaspoon",
		"nutmeg":"1/8 tablespoon",
		"lemon juice":"1 tablespoon"
	}`

	outputMap := map[string]string{}

	err := json.Unmarshal([]byte(jsonAppleReceipe), &outputMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(outputMap)
}
