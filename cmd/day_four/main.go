package main

import (
	"fmt"
	"os"

	matrices "github.com/avg-cs-student/advent-2024/pkg/matrices"
)

func main() {
	if len(os.Args) < 2 {
		panic("Expected PROG FILENAME")
	}

	filename := os.Args[1]
	rawContents, err := os.ReadFile(filename)
	if err != nil {
		panic("Expected to be able to read file contents")
	}

	solutionOne := solvePartOne(rawContents)
	fmt.Printf("Part one: %d\n", solutionOne)

	solutionTwo := solvePartTwo(rawContents)
	fmt.Printf("Part one: %d\n", solutionTwo)
}

func solvePartOne(rawContents []byte) int {
	matrix := matrices.ToRuneMatrix(rawContents)
	return check(matrix)
}

func solvePartTwo(rawContents []byte) int {
	matrix := matrices.ToRuneMatrix(rawContents)
	return check2(matrix)
}

func check(matrix [][]rune) int {
	xmas := [4]rune{'X', 'M', 'A', 'S'}
	total := 0

	var traverse func(int, int, int, func(int, int) (int, int)) int
	traverse = func(row, col, needed int, nextIndex func(int, int) (int, int)) int {
		if needed == len(xmas) {
			return 1
		}

		if row >= len(matrix) || row < 0 {
			return 0
		}

		if col >= len(matrix[row]) || col < 0 {
			return 0
		}

		if matrix[row][col] != xmas[needed] {
			return 0
		}

		nextRow, nextCol := nextIndex(row, col)
		return traverse(nextRow, nextCol, needed+1, nextIndex)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			upLeft := func(row, col int) (int, int) { return row - 1, col - 1 }
			upRight := func(row, col int) (int, int) { return row - 1, col + 1 }
			downLeft := func(row, col int) (int, int) { return row + 1, col - 1 }
			downRight := func(row, col int) (int, int) { return row + 1, col + 1 }
			forward := func(row, col int) (int, int) { return row, col + 1 }
			reverse := func(row, col int) (int, int) { return row, col - 1 }
			up := func(row, col int) (int, int) { return row - 1, col }
			down := func(row, col int) (int, int) { return row + 1, col }

			total += traverse(i, j, 0, upRight)
			total += traverse(i, j, 0, upLeft)
			total += traverse(i, j, 0, downRight)
			total += traverse(i, j, 0, downLeft)
			total += traverse(i, j, 0, forward)
			total += traverse(i, j, 0, reverse)
			total += traverse(i, j, 0, up)
			total += traverse(i, j, 0, down)
		}
	}

	return total
}

func check2(matrix [][]rune) int {
	total := 0

	checkMas := func(row, col int) int {
		sCount, mCount := 0, 0
		var side string
		if matrix[row-1][col-1] == 'S' {
			sCount++
			side = "top-left"
		}

		if matrix[row-1][col+1] == 'S' {
			sCount++
			side = "top-right"
		}

		if matrix[row+1][col+1] == 'S' {
			sCount++
			if side == "top-left" {
				return 0
			}
		}

		if matrix[row+1][col-1] == 'S' {
			sCount++
			if side == "top-right" {
				return 0
			}
		}

		if sCount != 2 {
			return 0
		}

		if matrix[row-1][col-1] == 'M' {
			mCount++
		}

		if matrix[row-1][col+1] == 'M' {
			mCount++
		}

		if matrix[row+1][col+1] == 'M' {
			mCount++
		}

		if matrix[row+1][col-1] == 'M' {
			mCount++
		}

		if mCount != 2 {
			return 0
		}

		return 1
	}

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix)-1; j++ {
			if matrix[i][j] != 'A' {
				continue
			}
			total += checkMas(i, j)
		}
	}

	return total
}
