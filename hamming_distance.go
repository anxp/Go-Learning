package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var s1 = "GAGCCTACTAACGGGAT"
	var s2 = "CATCGTAATGACGGCCT"
	var err error
	var hD int

	hD, err = hammingDistance(s1, s2)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Hamming distance = ", hD)
}

func hammingDistance(str1, str2 string) (int, error) {
	//Calculate and compare strings lengths. In symbols, NOT in bytes!
	strOneLength := len([]rune(str1))
	strTwoLength := len([]rune(str2))

	if strOneLength != strTwoLength {
			return 0, errors.New("Error! Strings have different length!")
	}

	//Proceed if strOneLength == strTwoLength:
	var numberOfSymbols = strOneLength

	var strOneSlice []rune
	var strTwoSlice []rune
	var hemmingDistance int

	//String to slice:
	for _, value := range str1 {
		strOneSlice = append(strOneSlice, value)
	}

	//String to slice:
	for _, value := range str2 {
		strTwoSlice = append(strTwoSlice, value)
	}

	for i:=0; i<numberOfSymbols; i++ {
		if strOneSlice[i] != strTwoSlice[i] {
			hemmingDistance++
		}
	}

	return hemmingDistance, nil
}
