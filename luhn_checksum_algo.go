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
	var cardNumberSlice []int //Here we will store card number digits
	var allDigitSum int = 0

	re := regexp.MustCompile("\\D") //Pattern for preg replace. Any symbol NOT a digit will be ignored
	filteredCardNumber = re.ReplaceAllString(cardNumber, "")

	//Let's make a slice from a string. It's SAFE to use len() here, because at our string ONLY digits are left!
	cardNumberSlice = make([]int, len(filteredCardNumber))

	//Fill our slice with digits:
	for key, value := range filteredCardNumber {
		cardNumberSlice[key] = int(value - '0') //WHAT A FUCK IS GOING ON?
	}

	//Reverse our slice from end to start:
	for i, j := 0, len(cardNumberSlice)-1; i < j; i,j = i+1, j-1 {
		cardNumberSlice[i], cardNumberSlice[j] = cardNumberSlice[j], cardNumberSlice[i]
	}

	for i:=0; i<len(cardNumberSlice); i++ {
		//We will now effectively iterate slice from end to start, and will take only EVEN elements:
		if i%2 !=0 { //Because EVEN elements has ODD indexes :)
			cardNumberSlice[i] = evaluateCurrentValue(cardNumberSlice[i])
		}

		allDigitSum += cardNumberSlice[i]
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
