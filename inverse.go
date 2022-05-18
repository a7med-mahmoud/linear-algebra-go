package main

import "errors"

// Inverse calculates the inverse of a matrix
func Inverse(mat Mat) (Mat, error) {
	det, err := Det(mat)
	if err != nil {
		return nil, err
	}

	if det == 0 {
		return nil, errors.New("can't calculate inverse of a matrix with determinant of 0")
	}

	trn, err := Transpose(mat)
	if err != nil {
		return nil, err
	}

	adj, err := Cofactors(trn)
	if err != nil {
		return nil, err
	}

	inv, err := ScalarMul(adj, 1/det)
	if err != nil {
		return nil, err
	}

	return inv, nil
}
