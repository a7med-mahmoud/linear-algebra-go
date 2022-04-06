package main

func Transpose(mat Mat) (Mat, error) {
	dim, err := GetDim(mat)
	if err != nil {
		return nil, err
	}
	trn := make(Mat, dim.Cols)
	for r, row := range mat {
		for c, num := range row {
			if r == 0 {
				trn[c] = make([]float64, dim.Rows)
			}
			trn[c][r] = num
		}
	}
	return trn, nil
}
