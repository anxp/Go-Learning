package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
	This function generates ONLY-DIGIT pass with given length. For example for length == 4 pass will be some random value in range [1000, 9999]
 */
func generateDigitPass(length int) func() map[int]int {
	passStorage := make(map[int]int)

	return func() map[int]int {
		var result int
		var maxPassValueForGivenLength int //999 for lengh == 3
		var minPassValueForGivenLength int //100 for length == 3

		maxPassValueForGivenLength = int (math.Pow10(length) - 1)
		minPassValueForGivenLength = int (math.Pow10(length - 1))

		//We need to initialize seed, or our random values will be predictable:
		rand.Seed(time.Now().UnixNano())

		result = rand.Intn(maxPassValueForGivenLength - minPassValueForGivenLength) + minPassValueForGivenLength

		passStorage[result]++

		return passStorage
	}
}

func main() {
	const maxNumberOfPasswords = 2500
	for i:=1; ;i++ { //We don't know how much iteration will be, so we do infinite loop, and break it when it generate valid data!
		allValuesAreUnique := true

		genPassPointer := generateDigitPass(i) //Remember pointer to function call in variable (genPassPointer)
		for j:=0; j < maxNumberOfPasswords - 1; j++ {
			genPassPointer() //Call THE SAME INSTANCE of function (maxNumberOfPasswords - 1) times...
		}
		
		resultMap := genPassPointer() //...and call it maxNumberOfPasswords-th time and get result.

		//Check if result map has duplicates or just unique values:
		for _, value := range resultMap {
			if value > 1 {
				allValuesAreUnique = false
				break
			}
		}

		if allValuesAreUnique {
			fmt.Println("Current password length is:", i, "No duplicates among", maxNumberOfPasswords, "generated passwords.")
			break
		} else {
			fmt.Println("Current password length is:", i, "Duplicates found, so we increase length +1 and will generate", maxNumberOfPasswords, "new passwords.")
		}
	}
}


