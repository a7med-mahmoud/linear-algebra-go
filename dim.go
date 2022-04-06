package main

import "errors"

// Dim returns the dimension of a matrix
// and checks if it's rectangular or not.
// the first return value is the number of rows,
// the second is the number of columns
// and the third is an error (if present)
func Dim(mat Mat) (uint8, uint8, error) {
	rows := len(mat[0])
	cols := len(mat)

	for _, row := range mat {
		if len(row) != rows {
			return 0, 0, errors.New("matrix is not rectangular")
		}
	}

	return uint8(rows), uint8(cols), nil
}
