/**
	Given a positive integer n> 1. Print all prime factors of this number in non-decreasing order.
	Complexity of algorithm should be O(logn).
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var valueToFactorize int
	var foundPrimeFactors []int

	fmt.Print("Enter an integer value, which should be factorized: ")
	fmt.Scan(&valueToFactorize)

	foundPrimeFactors = primeFactorsSearch(valueToFactorize)

	fmt.Println(foundPrimeFactors)
}

func primeFactorsSearch(valueToFactorize int, foundPrimeFactors ...[]int) []int {
	var maxPossibleFactor int
	var sqrtOfInputValue float64
	var alreadyFound []int

	// If foundPrimeFactors != nil -> means this is not first run of this function (recursive begun),
	// and slice with prime factors already has some values
	if foundPrimeFactors != nil {
		alreadyFound = foundPrimeFactors[0]
	}

	// We should devide valueToFactorize on these dividers (in cycle):
	// 2, 3, 4, 5, ... int(sqrt(valueToFactorize)), according to https://ru.wikipedia.org/wiki/Перебор_делителей
	// so at first, we need to calculate sqrt(valueToFactorize) and round this value to int:
	sqrtOfInputValue = math.Sqrt(float64(valueToFactorize))
	maxPossibleFactor = int(math.Round(sqrtOfInputValue))

	for i:=2; i <= maxPossibleFactor; i++ {
		if valueToFactorize % i == 0 {
			//Factor found! Let's add this devider to collection of previously found (if they are) factors:
			alreadyFound = append(alreadyFound, i)
			//... and run this function again, but this time with argument = result of division of initial value by found divider:
			return primeFactorsSearch(valueToFactorize / i, alreadyFound)
		}
	}

	// If we have not fall in recursion at this moment, -> means NO prime factors was found at CURRENT iteration ->
	// so CURRENT valueToFactorize is prime number, and CURRENT iteration is the last iteration, so just add it to collection of prime factors:
	alreadyFound = append(alreadyFound, valueToFactorize)

	//Begin return process from this most deep step (when final prime number found):
	return alreadyFound
}