package main

import (
	"fmt"
)

type Friend_Details struct {
	name           string
	mob            string
	alternate_Mob  string
	address        string
	city           string
	college_Friend bool
}

func main() {
	Friend1 := Friend_Details{
		name:           "Hemangi",
		mob:            "8989898989",
		alternate_Mob:  "7978787878",
		address:        "Bangalore, Karnataka",
		city:           "Bangalore",
		college_Friend: true,
	}
	Friend2 := Friend_Details{
		name:           "Geetha",
		mob:            "9009009009",
		alternate_Mob:  "-",
		address:        "Chennai Near XYZ ",
		city:           "Tamil Nadu",
		college_Friend: false,
	}
	Friend3 := Friend_Details{
		name:           "Sunitha",
		mob:            "8282828282",
		alternate_Mob:  "022543200",
		address:        "Mumbai",
		city:           "Maharashtra",
		college_Friend: true,
	}
	Friends_List := []Friend_Details{
		Friend1,
		Friend2,
		Friend3,
	}

	fmt.Println("College Friends are...")
	for _, friend := range Friends_List {
		if friend.college_Friend {
			fmt.Println(friend.name, ":", friend.city)
		}
	}

}
