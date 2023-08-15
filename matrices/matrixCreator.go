package matrices

import (
	"math/rand"
)

type Matrix struct {
	Length int
	Height int
	Matrix [][]int
}

func NewMatrix(length int, height int) *Matrix {
	m := &Matrix{
		Length: length,
		Height: height,
		Matrix: make([][]int, length, height),
	}
	return m
}

func (mat *Matrix) RandomizeMatrix() *Matrix {
	for l := 0; l < mat.Length; l++ {
		for h := 0; h < mat.Height; h++ {
			mat.Matrix[l] = append(mat.Matrix[l], rand.Intn(2*mat.Length)-mat.Length)
		}
	}
	return mat
}
