package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var hashesPool []string
	hashesCount := 1000000
	hashesPool = make([]string, hashesCount)

	startTime := time.Now()

	for i := 0; i < hashesCount; i++ {
		hashesPool[i] = randString(10)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Execution time is: ", elapsedTime, "sec.")
}

func randString(length int, additionalParams ...string) string {
	const defaultCharSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//const defaultCharSet = "日本語"
	//const defaultCharSet = "本日は入荷したての"
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

	return resultString
}
