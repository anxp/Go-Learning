package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePass(length int, additionalParams ...string) func() map[string]int {
	passStorage := make(map[string]int)

	return func() map[string]int {
		const defaultCharSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		var charSet string
		var charsSlice []rune
		var resultString string

		//Check if user specified string with characters or use default:
		if additionalParams != nil {
			charSet = additionalParams[0]
		} else {
			charSet = defaultCharSet
		}

		//We use for-range because only it correctly iterates multi-bytes characters:
		for _, value := range charSet {
			charsSlice = append(charsSlice, value)
		}

		//We need to initialize seed, or our random values will be predictable:
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < length; i++ {
			symbolIndex := rand.Intn(len(charsSlice))
			resultString += string(charsSlice[symbolIndex])
		}

		passStorage[resultString]++

		return passStorage
	}
}

func main() {
	for i:=1; ;i++ { //We don't know how much iteration will be, so we do infinite loop, and break it when it generate valid data!
		allValuesAreUnique := true

		genPassPointer := generatePass(i) //Remember pointer to function call in variable (genPassPointer)
		for j:=0; j < 299; j++ {
			genPassPointer() //Call THE SAME INSTANCE of function 299 times...
		}
		
		resultMap := genPassPointer() //...and call it 300-th time and get result.

		//Check if result map has duplicates or just unique values:
		for _, value := range resultMap {
			if value > 1 {
				allValuesAreUnique = false
				break
			}
		}

		if allValuesAreUnique {
			fmt.Println("Current password length is:", i, "No duplicates among 300 generated passwords.")
			break
		} else {
			fmt.Println("Current password length is:", i, "Duplicates found, so we increase length +1 and will generate 300 new passwords.")
		}
	}
}


