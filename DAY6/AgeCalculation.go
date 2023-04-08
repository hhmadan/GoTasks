package main

import (
	"fmt"
	"time"
)

type familyMember struct {
	name, relation    string
	date, month, year int
}

func main() {
	member1 := familyMember{
		name:     "Hemangi",
		relation: "Self",
		date:     22,
		month:    12,
		year:     2002,
	}
	member2 := familyMember{
		name:     "MotherName",
		relation: "Mother",
		date:     14,
		month:    04,
		year:     1972,
	}
	member3 := familyMember{
		name:     "FatherName",
		relation: "Father",
		date:     23,
		month:    04,
		year:     1970,
	}
	membersList := []familyMember{
		member1,
		member2,
		member3,
	}
	for _, eachMember := range membersList {
		computeAge(eachMember.name, eachMember.year)
	}
}

func computeAge(name string, birthYear int) {
	currentYear := time.Now().Year()
	ageOfMember := currentYear - birthYear
	fmt.Printf("Member==> Name: %s and Age: %d years.\n", name, ageOfMember)
}
