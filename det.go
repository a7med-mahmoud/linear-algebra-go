package main

import (
	"errors"
	"math"
)

func Det(mat Mat) (float64, error) {
	dim, err := GetDim(mat)

	if err != nil {
		return 0, err
	}

	if dim.Rows != dim.Cols {
		return 0, errors.New("can't calculate determinant of non-square matrix")
	}

	if dim.Rows == 2 {
		return mat[0][0]*mat[1][1] - mat[1][0]*mat[0][1], nil
	}

	det := 0.0
	for i, num := range mat[0] {
		sign := math.Pow(-1, float64(i))
		sub, err := Cancel(mat, 1, uint8(i)+1)
		if err != nil {
			return 0, err
		}
		subDet, err := Det(sub)
		if err != nil {
			return 0, err
		}
		det += sign * num * subDet
	}

	return det, nil
}
