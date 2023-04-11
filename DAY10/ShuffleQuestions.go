package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
		panic(e)
	}
}
func quizMenu(lines [][]string, shuffle bool) int {
	var totalQues int
	var quiz []QuesAndAns
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
		quiz = append(quiz, dataOfQuestion)

		if shuffle {
			r := rand.New(rand.NewSource(time.Now().Unix()))
			for n := len(quiz); n > 0; n-- {
				randIndex := r.Intn(n)
				quiz[n-1], quiz[randIndex] = quiz[randIndex], quiz[n-1]
			}
		}

		fmt.Println(dataOfQuestion.Question)
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("A.", dataOfQuestion.Option_a)
		fmt.Println("B.", dataOfQuestion.Option_b)
		fmt.Println("C.", dataOfQuestion.Option_c)
		fmt.Println("D.", dataOfQuestion.Option_d)
		fmt.Printf("Enter Your Answer: \n")
		response, err := reader.ReadString('\n')
		checkErr(err)
		fmt.Println("Your answer is: ", response, "Correct answer is:", dataOfQuestion.Correct_Option)
		processResult(response, dataOfQuestion.Correct_Option)
	}
	return totalQues, quiz
}

func processResult(response string, answer string) {
	res := strings.TrimSpace(response)
	ans := strings.TrimSpace(answer)
	if res == ans {
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
	shuffle := flag.Bool("shuffle", false, "Shuffle the order of questions")
	flag.Parse()
	filename := "Cricket_Players_Stats.csv"
	lines := processFile(&filename)
	var totalScore int
	totalScore = quizMenu(lines, *shuffle)
	fmt.Println("Number of correct questions is ", score, "Total question is", totalScore)
}
