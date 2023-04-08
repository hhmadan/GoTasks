package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Friend_Details struct {
	name          string
	mob           string
	alternate_Mob string
	address       string
	city          string
}

func main() {
	Friend1 := Friend_Details{
		name:          "Hemangi",
		mob:           "8989898989",
		alternate_Mob: "7978787878",
		address:       "Bangalore, Karnataka",
		city:          "Bangalore",
	}
	Friend2 := Friend_Details{
		name:          "Geetha",
		mob:           "9009009009",
		alternate_Mob: "-",
		address:       "Chennai Near XYZ ",
		city:          "Tamil Nadu",
	}
	Friends_List := []Friend_Details{
		Friend1,
		Friend2,
	}
	// for _, friend := range Friends_List {
	// 	fmt.Println(friend.name, " : ", friend.mob)
	// }
	file, err := os.Create("Friend_Records.csv")
	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	for _, record := range Friends_List {
		row := []string{record.name, record.mob, record.alternate_Mob, record.address, record.city}

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
