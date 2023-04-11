package main

import (
	"bufio"
	"fmt"
	"os"
)

// Use os.open() function to open the file.
// Use bufio.NewScanner() function to create the file scanner.
// Use bufio.ScanLines() function with the scanner to split the file into lines.
// Then use the scanner Scan() function in a for loop to get each line and process it.
func main() {

	readFile, err := os.Open("LearnerData.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScannerLines := bufio.NewScanner(readFile)

	fileScannerLines.Split(bufio.ScanLines)

	fileScannerWords := bufio.NewScanner(readFile)
	fileScannerWords.Split(bufio.ScanWords)

	for fileScannerLines.Scan() {
		fmt.Println("Printing New Line: ", fileScannerLines.Text())

	}
}
