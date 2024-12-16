package matrices

import (
	"fmt"
)

func ToRuneMatrix(data []byte) [][]rune {
	var matrix [][]rune
	matrix = append(matrix, []rune{})

	row := 0
	for i, b := range data {
		switch b {
		case '\n':
			row++
			if i == len(data)-1 {
				break
			}
			matrix = append(matrix, []rune{})
		default:
			matrix[row] = append(matrix[row], rune(b))
		}
	}

	return matrix
}

func DisplayMatrix(matrix [][]rune) {
	for i := range matrix {
		if i == 0 {
			fmt.Printf("%2c ", ' ')
			for j := range matrix[0] {
				fmt.Printf("%2d ", j)
			}
			fmt.Println()
		}
		for j := range matrix[i] {
			if j == 0 {
				fmt.Printf("%2d ", i)
			}
			fmt.Printf("%2c ", matrix[i][j])
		}
		fmt.Println()
	}
}
