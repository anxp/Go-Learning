package main

import (
	"fmt"
	"regexp"
)

func main() {
	result := luhnChecksum("4539 1488 0343 6467 本日は入荷したての")

	if (result) {
		fmt.Println("Card number passed verification!")
	} else {
		fmt.Println("Card number not valid!")
	}
}

func luhnChecksum(cardNumber string) bool {
	var filteredCardNumber string //Here we will store clean card number, without spaces and letters
	var allDigitSum int

	re := regexp.MustCompile("\\D") //Pattern for preg replace. Any symbol NOT a digit will be ignored
	filteredCardNumber = re.ReplaceAllString(cardNumber, "")

	//Here we moving on string, but from end to start:
	for i:=len(filteredCardNumber) - 2; i >= 0; i -= 2 {
		evenFromEndValue := int(filteredCardNumber[i] - '0')
		oddFromEndValue := int(filteredCardNumber[i+1] - '0')

		allDigitSum += evaluateCurrentValue(evenFromEndValue) + oddFromEndValue
	}

	if allDigitSum%10 == 0 {
		return true
	} else {
		return false
	}
}

func evaluateCurrentValue(currentValue int) int {
	if tmp := currentValue*2; tmp <= 9 {
		return tmp
	} else {
		return tmp - 9
	}
}
