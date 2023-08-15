package main

import (
	"fmt"

	"github.com/NicoleKaldus/m/matrices"
)

func main() {

	mat1 := matrices.NewMatrix(4, 4)
	fmt.Println((mat1.RandomizeMatrix()))

	mat2 := matrices.NewMatrix(4, 4)
	fmt.Println((mat2.RandomizeMatrix()))

	// fmt.Println(matrices.AddMatrices(mat1, mat2))
	// fmt.Println(matrices.SubtractMatrices(mat1, mat2))
	fmt.Println(matrices.MultiplyMatricies(mat1, mat2))
}
