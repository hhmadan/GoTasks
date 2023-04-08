package main

import (
	"encoding/json"
	"fmt"
)

type Friend_Details struct {
	Name           string
	Mob            string
	Alternate_Mob  string
	Address        string
	City           string
	College_Friend bool
}

func main() {
	Friend1 := Friend_Details{
		Name:           "Hemangi",
		Mob:            "8989898989",
		Alternate_Mob:  "7978787878",
		Address:        "Bangalore, Karnataka",
		City:           "Bangalore",
		College_Friend: true,
	}
	Friend2 := Friend_Details{
		Name:           "Geetha",
		Mob:            "9009009009",
		Alternate_Mob:  "-",
		Address:        "Chennai Near XYZ ",
		City:           "Tamil Nadu",
		College_Friend: false,
	}
	Friend3 := Friend_Details{
		Name:           "Sunitha",
		Mob:            "8282828282",
		Alternate_Mob:  "022543200",
		Address:        "Mumbai",
		City:           "Maharashtra",
		College_Friend: true,
	}
	Friends_List := []Friend_Details{
		Friend1,
		Friend2,
		Friend3,
	}

	//fmt.Println(Friends_List)

	bytes, _ := json.MarshalIndent(Friends_List, " ", "\t")
	fmt.Println(string(bytes))

	jsonFile := string(bytes)

	byteString := []byte(jsonFile)
	var friend []Friend_Details
	err := json.Unmarshal(byteString, &friend)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("------JSON FILE TO SLICE OF STRUCT------")
	fmt.Println(friend)
}
