package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getAverage(numbersList []float64) float64 {
	sum := 0.0

	for _, num := range numbersList {
		sum += num
	}

	return sum / float64(len(numbersList))
}

func getStandardDeviation(numbersList []float64) float64 {
	sumOfSquares := 0.0

	average := getAverage(numbersList)

	for _, num := range numbersList {
		sumOfSquares += (num - average) * (num - average)
	}

	return math.Sqrt(sumOfSquares / float64(len(numbersList)))
}

func main() {
	var numbersList []float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a list of floating point numbers...If need to exit enter 'q' or 'Q' to stop:")
	for scanner.Scan() {
		inputNum := scanner.Text()

		if strings.ToLower(inputNum) == "q" {
			break
		}

		num, err := strconv.ParseFloat(inputNum, 64)
		if err != nil {
			fmt.Println("Invalid Data")
			continue
		}
		numbersList = append(numbersList, num)
	}

	average := getAverage(numbersList)
	standardDeviation := getStandardDeviation(numbersList)
	fmt.Printf("Average: %f\n", average)
	fmt.Printf("Standard deviation: %f\n", standardDeviation)
}
