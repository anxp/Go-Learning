package main

import (
	"fmt"
	"sort"
)

func main() {
	var maxValue int
	var primeNumbersUnderMaxValue []int

	fmt.Print("Enter max number, under which we will search prime numbers: ")
	fmt.Scan(&maxValue)

	primeNumbersUnderMaxValue = eratostheneSieve(maxValue)

	fmt.Println(primeNumbersUnderMaxValue)
}

func eratostheneSieve(maxValue int) []int {
	var maxSieveIndex int
	var primeNumbersCollection []int

	if even(maxValue) {
		maxSieveIndex = maxValue - 1
	} else {
		maxSieveIndex = maxValue
	}

	// Initialization of Eratosthene Sieve. Keys of map are numbers, which will be checked for simplicity.
	// Values of map are 1 or 0 - is number represented by key is prime or not.
	// For better memory performance, we skip all even values except 2 (because it's obviously prime).
	sieve := make(map[int]int)
	sieve[2] = 1
	for i:=3; i<=maxSieveIndex; i+=2 {
		sieve[i] = 1
	}

	for j:=3; j<=maxSieveIndex; j+=2 {
		//If this value is already deleted (composite number detected), just skip it:
		if sieve[j] == 0 {
			continue
		}

		//No sense to continue algo. At this point all composite numbers up to maxSieveIndex already deleted, so break the loop:
		if j*j > maxSieveIndex {
			break
		}

		//The essence of the Eratosthene algorithm ->
		//Let's delete all values, of multiple j, starting from k, but except k (and repeat this many times by the outer loop):
		for k:=j; k<=maxSieveIndex; k+=2 {
			if k == j {
				continue
			}

			if k%j == 0 {
				sieve[k] = 0
			}
		}
	}

	//Copy all found prime numbers to slice:
	for key, value := range sieve {
		if value != 0 {
			primeNumbersCollection = append(primeNumbersCollection, key)
		}
	}

	//And, finally, sort:
	sort.Ints(primeNumbersCollection)

	return primeNumbersCollection
}

func even(number int) bool {
	return number%2 == 0
}