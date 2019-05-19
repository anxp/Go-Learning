package main

import "golang.org/x/tour/pic"

func picTroughRange(dx, dy int) [][]uint8 {

	var matrix2D [][]uint8
	matrix2D = make([][]uint8, dy) //Allocate "outer" container

	for outerKey := range matrix2D {
		matrix2D[outerKey] = make([]uint8, dx)
		for innerKey := range matrix2D[outerKey] {
			matrix2D[outerKey][innerKey] = uint8((innerKey + outerKey) / 2)
		}
	}

	return matrix2D
}

func picTroughFor(dx, dy int) [][]uint8 {
	var matrix2D [][]uint8
	matrix2D = make([][]uint8, dy) //Allocate "outer" container

	for j := 0; j < dy; j++ {
		matrix2D[j] = make([]uint8, dx)
		for i := 0; i < dx; i++ {
			matrix2D[j][i] = uint8(i * j)
		}
	}

	return matrix2D
}

func main() {
	pic.Show(picTroughRange)
	//pic.Show(picTroughFor)
}
