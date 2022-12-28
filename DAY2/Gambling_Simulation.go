package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gamblingSimulation(stake int, goal int, trial int) {
	totalwin := 0
	totalbets := 0
	for i := 0; i < trial; i++ {
		bets := 0
		currentstake := stake
		for currentstake > 0 && currentstake < goal {
			bets++
			rand.Seed(time.Now().UnixNano())
			outcome := rand.Intn(2) //to win=1 or loss=0
			if outcome == 1 {
				currentstake++
			} else {
				currentstake--
			}
		}
		if currentstake == goal {
			totalwin++
		}
		totalbets += bets
	}

	winPercentage := float64(totalwin) / float64(trial) * 100
	winAvg := float64(totalwin) / float64(trial)

	fmt.Println("Number of wins: ", totalwin)
	fmt.Println("Number of bets: ", totalbets)
	fmt.Println("Win Percentage: ", winPercentage)
	fmt.Println("Average Win: ", winAvg)

}
func main() {
	var stake, goal, trial int
	fmt.Println("Enter Stake: ")
	fmt.Scan(&stake)
	fmt.Println("Enter Goal: ")
	fmt.Scan(&goal)
	fmt.Println("Enter Number of Trials: ")
	fmt.Scan(&trial)
	gamblingSimulation(stake, goal, trial)
}
