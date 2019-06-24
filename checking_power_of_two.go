package main

import "fmt"

func main() {
	var checkingValue int

	fmt.Print("Enter a integer value, you want to check for power of TWO:  ")
	fmt.Scan(&checkingValue)

	if checkPowerOfTwo(checkingValue) {
		fmt.Print("This value is exact power of two.")
	} else {
		fmt.Print("This value is NOT the power of TWO!")
	}
}

func checkPowerOfTwo(inputValue int) bool {
	//Set up exit condition:
	if inputValue == 1 {
		return true
	}

	//If inputValue devided completely on 2, call function recursively,
	//else -> return false, as given value is not a power of two:
	if inputValue%2 == 0 {
		intermediateValue := inputValue / 2
		return checkPowerOfTwo(intermediateValue)
	} else {
		return false
	}
}
