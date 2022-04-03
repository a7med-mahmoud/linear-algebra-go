package main

import "fmt"

type Mat [][]float64

func main() {
	fmt.Println("Linear Algebra")

	mat := Mat{
		{1, 2, 3, 24},
		{4, 2, 12, 7},
		{7, -1, 9, 2},
		{7, -1, 9, 12},
	}

	det, err := Det(mat)
	if err != nil {
		panic(err)
	}

	fmt.Println(det)
}
