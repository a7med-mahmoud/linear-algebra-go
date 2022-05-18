package main

import (
	"errors"
	"math"
)

// Cofactors calculates matrix of confactors of a matirx
func Cofactors(mat Mat) (Mat, error) {
	rows, cols, err := Dim(mat)
	if err != nil {
		return nil, err
	}

	if rows != cols {
		return nil, errors.New("can't calculate matrix of cofactors of non-square matrix")
	}

	cof := make(Mat, rows)

	for r, row := range mat {
		for c := range row {
			if c == 0 {
				cof[r] = make([]float64, cols)
			}

			sign := math.Pow(-1, float64(r+c+2))
			sub, err := Cancel(mat, uint8(r)+1, uint8(c)+1)
			if err != nil {
				return nil, err
			}

			det, err := Det(sub)
			if err != nil {
				return nil, err
			}

			cof[r][c] = sign * det
		}
	}

	return cof, nil
}
