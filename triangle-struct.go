package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

//Struct, which describes triangle properties. Our triangle defined by three sides.
type Triangle struct {
	sideA, sideB, sideC float64
}

//This is a setter for all 3 triangle sides
func (t *Triangle) SetSides(a, b, c float64) {
	t.sideA = a
	t.sideB = b
	t.sideC = c
}

//Method to check if triangle right or not (right triangle has 90-degree angle)
func (t Triangle) IsThisRightTriangle() bool {
	if math.Pow(t.sideA, 2)+math.Pow(t.sideB, 2) == math.Pow(t.sideC, 2) {
		return true
	} else {
		return false
	}
}

//Method to calculate triangle area by Heron's formula
func (t Triangle) GetArea() (float64, error) {
	var halfPerimeter float64
	var area float64
	var err error

	if t.isThisATriangle() {
		halfPerimeter = (t.sideA + t.sideB + t.sideC) / 2
		area = math.Sqrt(halfPerimeter * (halfPerimeter - t.sideA) * (halfPerimeter - t.sideB) * (halfPerimeter - t.sideC))
		return area, err
	} else {
		return 0, errors.New("Submitted data (sides) does not compose in triangle!")
	}
}

//Method to calculate area of _right_ triangle. Just as (a*b)/2
func (t Triangle) GetRightTriangleArea() (float64, error) {
	var err error
	if t.IsThisRightTriangle() {
		return (t.sideA * t.sideB) / 2, err
	} else {
		return 0, errors.New("Can't apply GetRightTriangleArea() method to non-right triangle!")
	}

}

//Just _most_simplest_ implementation of Bubble Sorting Algo:
func (t Triangle) sortDesc(x, y, z float64) (float64, float64, float64) {
	var inputData [3]float64 = [3]float64{x, y, z}

	for j := 0; j < len(inputData)-1; j++ {
		for i := 0; i < len(inputData)-1; i++ {
			if inputData[i] < inputData[i+1] {
				inputData[i], inputData[i+1] = inputData[i+1], inputData[i]
			}
		}
	}

	return inputData[0], inputData[1], inputData[2]
}

//Method to check if user submitted 3 sides can make a triangle or not
func (t Triangle) isThisATriangle() bool {
	var a, b, c = t.sortDesc(t.sideA, t.sideB, t.sideC)

	if a < (b + c) {
		return true
	} else {
		return false
	}
}

func main() {
	var sideA, sideB, sideC, area float64
	var ourTriangle Triangle
	var err error

	fmt.Println("This program calculates triangle area and makes check if triangle is right or not")
	fmt.Print("Enter size of side A: ")
	fmt.Scan(&sideA)

	fmt.Print("Enter size of side B: ")
	fmt.Scan(&sideB)

	fmt.Print("Enter size of side C: ")
	fmt.Scan(&sideC)

	ourTriangle.SetSides(sideA, sideB, sideC)

	fmt.Printf("Your input data is: side A = %f, side B = %f, side C = %f\n", ourTriangle.sideA, ourTriangle.sideB, ourTriangle.sideC)

	if ourTriangle.IsThisRightTriangle() {
		fmt.Println("Submited sides makes a right triangle.")
		area, err = ourTriangle.GetRightTriangleArea()
	} else {
		fmt.Println("Your triangle is not right. Let's check if this is a regular triangle, or not a triangle at all...")
		area, err = ourTriangle.GetArea()
		if err == nil {
			fmt.Println("This is a regular, but not right triangle.")
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("Calculated area of triangle: ", area)
	}
}
