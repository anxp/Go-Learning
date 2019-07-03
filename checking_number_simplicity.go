package main

import "fmt"

func main() {
	var checkingValue int

	fmt.Print("Enter an integer value, you want to check for simplicity: ")
	fmt.Scan(&checkingValue)

	if primeCheck(checkingValue) {
		fmt.Print("This number is prime.")
	} else {
		fmt.Print("This number is NOT prime.")
	}
}

func primeCheck(inputValue int, additionalParams ...int) bool {
	var devider int

	if additionalParams != nil {
		devider = additionalParams[0] //all next recursive iterations - we decrease devider by 1 and call checkNumberSimplicity again and again
	} else {
		devider = inputValue - 1 //devider initialization on first run
	}

	//We checked all devider values from inputValue-1 to 1 and didn't find any "integer division",
	// so we conclude, that checking number is prime.
	if devider == 1 {
		return true
	}

	if inputValue%devider == 0 {
		return false
	} else {
		return primeCheck(inputValue, devider - 1)
	}
}
