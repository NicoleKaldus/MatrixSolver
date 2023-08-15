package matrices

import (
	"errors"
)

func matArithmetic(addOrSubtract string, mat1 *Matrix, mat2 *Matrix) (*Matrix, error) {
	if mat1.Height != mat2.Height || mat1.Length != mat2.Length {
		return nil, errors.New("Unrecognized arithmetic operation.")
	}

	if addOrSubtract != "add" && addOrSubtract != "subtract" {
		return nil, errors.New("Matrices must be the same size in order to " + addOrSubtract + ".")
	}

	resMat := NewMatrix(mat1.Length, mat1.Height)

	for l := 0; l < resMat.Length; l++ {
		for h := 0; h < resMat.Height; h++ {
			res := 0
			if addOrSubtract == "add" {
				res = mat1.Matrix[l][h] + mat2.Matrix[l][h]
			} else {
				res = mat1.Matrix[l][h] - mat2.Matrix[l][h]
			}
			resMat.Matrix[l] = append(resMat.Matrix[l], res)
		}
	}

	return resMat, nil
}

func AddMatrices(mat1 *Matrix, mat2 *Matrix) (*Matrix, error) {
	return matArithmetic("add", mat1, mat2)
}

func SubtractMatrices(mat1 *Matrix, mat2 *Matrix) (*Matrix, error) {
	return matArithmetic("subtract", mat1, mat2)
}

func MultiplyMatricies(mat1 *Matrix, mat2 *Matrix) (*Matrix, error) {
	if mat1.Length != mat2.Height {
		return nil, errors.New("Length of matrix 1 and height of matrix 2 must be equal in order to multiply.")
	}

	resMat := NewMatrix(mat2.Length, mat1.Height)

	calculateTile := func(row *[]int, multiplyingMat *Matrix, colNum int) int {
		/*

			a b c		j k l		aj+bm+cp  ak+bn+cq
			d e f	x	m n o	=
			g h i		p q r

		*/
		resNum := 0
		for i := 0; i < len(*row); i++ {
			resNum += (*row)[i] * multiplyingMat.Matrix[i][colNum]
		}
		return resNum
	}

	for l := 0; l < resMat.Length; l++ {
		for h := 0; h < resMat.Height; h++ {
			resMat.Matrix[l] = append(resMat.Matrix[l], calculateTile(&mat1.Matrix[l], mat2, h))
		}
	}

	return resMat, nil
}
