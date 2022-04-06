package main

func Transpose(mat Mat) (Mat, error) {
	rows, cols, err := Dim(mat)
	if err != nil {
		return nil, err
	}
	trn := make(Mat, cols)
	for r, row := range mat {
		for c, num := range row {
			if r == 0 {
				trn[c] = make([]float64, rows)
			}
			trn[c][r] = num
		}
	}
	return trn, nil
}
