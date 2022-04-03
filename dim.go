package main

import (
	"errors"
	"fmt"
)

type Dim struct {
	Rows, Cols uint8
}

func NewDim(rows, cols int) *Dim {
	return &Dim{
		Rows: uint8(rows),
		Cols: uint8(cols),
	}
}

func (d *Dim) String() string {
	return fmt.Sprint(d.Rows, "x", d.Cols)
}

func GetDim(mat [][]float64) (*Dim, error) {
	rows := len(mat[0])
	cols := len(mat)

	for _, row := range mat {
		if len(row) != rows {
			return nil, errors.New("matrix is not rectangular")
		}
	}

	return NewDim(rows, cols), nil
}
