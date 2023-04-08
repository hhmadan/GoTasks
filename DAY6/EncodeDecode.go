package main

import "fmt"

func encodeString(inputString string) []byte {
	var byteDecodedString []byte
	for word := range inputString {
		decodedString := ((inputString[word]) + 1) //converted to string to (byte+1) eg. H: 72 ==>73
		byteDecodedString = append(byteDecodedString, decodedString)
	}
	return byteDecodedString
}

func decodeString(byteDecodedString []byte) string {
	var byteEndcodedString []byte
	for word := range byteDecodedString {
		endcodedString := ((byteDecodedString[word]) - 1)
		byteEndcodedString = append(byteEndcodedString, endcodedString)
	}
	return string(byteEndcodedString)
}

func main() {
	inputString := "Golang"
	xyz := encodeString(inputString)
	fmt.Println("Encoded String: ", string(xyz))
	abc := decodeString(xyz)
	fmt.Println("Decoded String: ", abc)
}
