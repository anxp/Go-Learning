package main

import (
	"fmt"
)

func create3Dmatrix(dx, dy, dz int) [][][]float64 {
	var matrix3D [][][]float64
	matrix3D = make([][][]float64, dz)

	for zIndex := range matrix3D {
		matrix3D[zIndex] = make([][]float64, dy)
		for yIndex := range matrix3D[zIndex] {
			matrix3D[zIndex][yIndex] = make([]float64, dx)
			for xIndex := range matrix3D[zIndex][yIndex] {
				matrix3D[zIndex][yIndex][xIndex] = float64(xIndex+yIndex+zIndex) / 2
			}
		}
	}

	return matrix3D
}

func printMatrix(matrix [][][]float64) {
	for zIndex, zValue := range matrix {
		//fmt.Println("Current zIndex: ", zIndex)
		for yIndex, yValue := range zValue {
			//fmt.Println("Current yIndex: ", yIndex)
			for xIndex, xValue := range yValue {
				fmt.Printf("Element[x=%d][y=%d][z=%d]: %f\n", xIndex, yIndex, zIndex, xValue)
				//fmt.Println("Element[x=", xIndex, "][y=", yIndex, "][z=", zIndex, "]: ", xValue)
			}
		}
	}
}

func main() {
	var matrix [][][]float64
	matrix = create3Dmatrix(3, 3, 3)
	printMatrix(matrix)
}
