package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var arg1, arg2, result float64
	var err error
	var operator string

	fmt.Print("Enter first argument: ")
	_,err = fmt.Scan(&arg1)

	exitOnError(err)

	fmt.Print("Enter second argument: ")
	_,err = fmt.Scan(&arg2)

	exitOnError(err)

	fmt.Print("Enter operator ([+][-][*][/]): ")
	_,err = fmt.Scan(&operator)

	exitOnError(err)

	result,err = calculate(arg1, arg2, operator)

	exitOnError(err)

	fmt.Printf("%f %s %f = %f", arg1, operator, arg2, result)
}

func calculate(arg1, arg2 float64, op string) (float64, error) {
	var result float64
	var err error

	switch op {
	case "+":
		result = arg1 + arg2
	case "-":
		result = arg1 - arg2
	case "*":
		result = arg1 * arg2
	case "/":
		if arg2 == 0 {
			return 0, errors.New("Error: Division by 0!")
		}
		result = arg1 / arg2
	default:
		return 0, errors.New("Invalid operation!")
	}

	return result, err
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}