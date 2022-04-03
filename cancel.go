package main

import (
	"fmt"
)

// Cancel is a function that removes a row and a column
// from a matrix and returns the resulting matrix.
// This is useful when calculating determinants for example.
func Cancel(mat Mat, row, col uint8) (Mat, error) {
	dim, err := GetDim(mat)
	if err != nil {
		return nil, err
	}
	if row > dim.Rows {
		return nil, fmt.Errorf("can't cancel row #%d from %d rows", row, dim.Rows)
	}
	if col > dim.Cols {
		return nil, fmt.Errorf("can't cancel column #%d from %d columns", row, dim.Cols)
	}
	ans := Mat{}
	for r, nums := range mat {
		if uint8(r) == row-1 {
			continue
		}

		curr := []float64{}
		for c, num := range nums {
			if uint8(c) != col-1 {
				curr = append(curr, num)
			}
		}
		ans = append(ans, curr)
	}
	return ans, nil
}
