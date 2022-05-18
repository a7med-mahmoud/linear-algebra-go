package main

// ScalarMul multiplies a matrix by a scalar
func ScalarMul(mat Mat, scalar float64) (Mat, error) {
	rows, cols, err := Dim(mat)
	if err != nil {
		return nil, err
	}

	ans := make(Mat, rows)

	for r, row := range mat {
		for c, num := range row {
			if c == 0 {
				ans[r] = make([]float64, cols)
			}

			ans[r][c] = num * scalar
		}
	}

	return ans, nil
}
