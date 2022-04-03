package main

import "fmt"

type Mat [][]float64

func main() {
	fmt.Println("Linear Algebra")

	mat := Mat{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	dim, err := GetDim(mat)
	if err != nil {
		panic(err)
	}

	fmt.Println(dim)
}
