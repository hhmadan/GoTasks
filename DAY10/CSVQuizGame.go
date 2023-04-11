package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type QuesAndAns struct {
	Question       string
	Option_a       string
	Option_b       string
	Option_c       string
	Option_d       string
	Correct_Option string
}

var score int

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func quizMenu(lines [][]string) int {
	var totalQues int
	for _, line := range lines {
		totalQues++
		dataOfQuestion := QuesAndAns{
			Question:       line[0],
			Option_a:       line[1],
			Option_b:       line[2],
			Option_c:       line[3],
			Option_d:       line[4],
			Correct_Option: line[5],
		}

		fmt.Println(dataOfQuestion.Question)
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("A.", dataOfQuestion.Option_a)
		fmt.Println("B.", dataOfQuestion.Option_b)
		fmt.Println("C.", dataOfQuestion.Option_c)
		fmt.Println("D.", dataOfQuestion.Option_d)
		fmt.Printf("Enter Your Response: \n")
		response, err := reader.ReadString('\n')
		checkErr(err)
		fmt.Println("Your answer is: ", response, "Correct answer is:", dataOfQuestion.Correct_Option)
		processResult(response, dataOfQuestion.Correct_Option)
	}
	return totalQues
}

func processResult(response string, answer string) {
	res := strings.TrimSpace(response)
	ans := strings.TrimSpace(answer)
	res1 := strings.ToUpper(res)
	if res1 == ans {
		score++
	}
}

func processFile(filename *string) [][]string {
	file, err := os.Open(*filename)
	checkErr(err)

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	checkErr(err)
	return lines
}

func main() {
	filename := "Cricket_Players_Stats.csv"
	lines := processFile(&filename)
	var totalScore int
	totalScore = quizMenu(lines)
	fmt.Println("Number of correct questions is ", score, "Total question is", totalScore)
}
