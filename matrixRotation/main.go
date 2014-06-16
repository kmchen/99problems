package main

import (
	"fmt"
)

// Strategy : Rotate row by row, copy must be done counter-clockwise
// Left -> Top, Bottom -> Left, Right -> Bottom, Top -> Right
func rotate(matrix [][]int, n int) {

	bound := n - 1
	// Rotate row by row
	for i := 0; i < n/2; i++ {
		for j := 0; j < bound-i; j++ {
			top := matrix[i][j+i]

			// left -> top
			matrix[i][j+i] = matrix[bound-j][i]

			// bottom -> left
			matrix[bound-j][i] = matrix[bound][bound-j]

			// right -> bottom
			matrix[bound][bound-j] = matrix[j+i][bound]

			// top -> right
			matrix[i+j][bound] = top
		}
		fmt.Println()
		//[[13 9 5 1] [14 6 7 2] [15 10 11 3] [16 12 8 4]]
		bound--
	}
	for row := 0; row <= n-1; row++ {
		for col := 0; col <= n-1; col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	rotate(matrix, 4)
}
