package main

import (
	"fmt"
	"os"

	matrices "github.com/avg-cs-student/advent-2024/pkg/matrices"
)

func main() {
	if len(os.Args) != 2 {
		panic("Expected PROG FILENAME")
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Expected to be able to read file.")
	}

	matrix := matrices.ToRuneMatrix(data)

	// solutionOne := solvePartOne(matrix)
	// fmt.Printf("Part one: %d\n", solutionOne)
	solutionTwo := solvePartTwo(matrix)
	fmt.Printf("Part two: %d\n", solutionTwo)
}

func solvePartOne(matrix [][]rune) int {
	maxRow, maxCol := len(matrix), len(matrix[0])

	uniq := 1
	curPos, facingIndex, err := findGuard(matrix)
	if err != nil {
		panic(err)
	}
	prevPos := &position{row: curPos.row, col: curPos.col}

	for curPos.row < maxRow && curPos.col < maxCol {
		if collided(matrix, curPos) {
			facingIndex = (facingIndex + 1) % len(guardRunes)
			curPos = prevPos
		}
		if matrix[curPos.row][curPos.col] == '.' {
			matrix[curPos.row][curPos.col] = 'X'
			uniq++
		}

		prevPos = curPos
		facing := guardRunes[facingIndex]
		curPos = directionFns[facing](prevPos)
	}

	return uniq
}

type position struct {
	row int
	col int
}

func collided(matrix [][]rune, pos *position) bool {
	return matrix[pos.row][pos.col] == '#'
}

type walkFn func(*position) *position

func walkUp(initial *position) *position {
	newPosition := &position{
		row: initial.row - 1,
		col: initial.col,
	}
	return newPosition
}

func walkDown(initial *position) *position {
	newPosition := &position{
		row: initial.row + 1,
		col: initial.col,
	}
	return newPosition
}

func walkRight(initial *position) *position {
	newPosition := &position{
		row: initial.row,
		col: initial.col + 1,
	}
	return newPosition
}

func walkLeft(initial *position) *position {
	newPosition := &position{
		row: initial.row,
		col: initial.col - 1,
	}
	return newPosition
}

var (
	guardRunes   = []rune{'^', '>', 'v', '<'}
	directionFns = map[rune]walkFn{
		guardRunes[0]: walkUp,
		guardRunes[1]: walkRight,
		guardRunes[2]: walkDown,
		guardRunes[3]: walkLeft,
	}
)

func findGuard(matrix [][]rune) (*position, int, error) {
	for i := range matrix {
		for j := range matrix[i] {
			for k, facing := range guardRunes {
				if matrix[i][j] == facing {
					return &position{row: i, col: j}, k, nil
				}
			}
		}
	}

	return nil, -1, fmt.Errorf("Expected to find a guard in the matrix.")
}

func solvePartTwo(matrix [][]rune) int {
	// TODO: debug only
	debugView := make([][]rune, len(matrix))
	for i := range debugView {
		debugView[i] = make([]rune, len(matrix[i]))
		copy(debugView[i], matrix[i])
	}

	maxRow, maxCol := len(matrix), len(matrix[0])

	uniq := 0
	curPos, facingIndex, err := findGuard(matrix)
	if err != nil {
		panic(err)
	}

	isOnTheMap := func(p *position) bool {
		return p.row < maxRow && p.row >= 0 && p.col < maxCol && p.col >= 0
	}

	curFacing := guardRunes[facingIndex]
	for {
		debugView[curPos.row][curPos.col] = curFacing
		matrix[curPos.row][curPos.col] = curFacing

		nextPos := directionFns[curFacing](curPos)
		if !isOnTheMap(nextPos) {
			break
		}

		for getRune(matrix, nextPos) == '#' {
			facingIndex = (facingIndex + 1) % len(guardRunes)
			curFacing = guardRunes[facingIndex]
			nextPos = directionFns[curFacing](curPos)
		}

		colContains := func(m [][]rune, target rune, rowStart, rowStop, col int) int {
			if rowStart < 0 || rowStop > maxRow {
				return -1
			}

			for i := rowStart; i < rowStop; i++ {
				if m[i][col] == target {
					return i
				}
			}

			return -1
		}

		rowContains := func(m [][]rune, target rune, colStart, colStop, row int) int {
			if colStart < 0 || colStop > maxCol {
				return -1
			}

			for i := colStart; i < colStop; i++ {
				if m[row][i] == target {
					return i
				}
			}

			return -1
		}

		switch curFacing {
		case '^':
			if rowContains(matrix, '>', curPos.col, maxCol, curPos.row) >= 0 {
				uniq++
			}
			revDirIndex := colContains(matrix, 'v', curPos.row+1, maxRow, curPos.col)
			blockerIndex := colContains(matrix, '#', curPos.row+1, maxRow, curPos.col)
			if blockerIndex >= 0 && revDirIndex >= 0 && revDirIndex < blockerIndex {
				uniq++
			}
		case 'v':
			if rowContains(matrix, '<', 0, curPos.col, curPos.row) >= 0 {
				uniq++
			}
			revDirIndex := colContains(matrix, '^', 0, curPos.row, curPos.col)
			blockerIndex := colContains(matrix, '#', 0, curPos.row, curPos.col)
			if blockerIndex >= 0 && revDirIndex >= 0 && revDirIndex > blockerIndex {
				uniq++
			}
		case '>':
			if colContains(matrix, 'v', curPos.row, maxRow, curPos.col) >= 0 {
				uniq++
			}
			revDirIndex := rowContains(matrix, '<', 0, curPos.col, curPos.row)
			blockerIndex := rowContains(matrix, '#', 0, curPos.col, curPos.row)
			if blockerIndex >= 0 && revDirIndex >= 0 && revDirIndex > blockerIndex {
				uniq++
			}
		case '<':
			if colContains(matrix, '^', 0, curPos.row, curPos.col) >= 0 {
				uniq++
			}
			revDirIndex := rowContains(matrix, '<', curPos.col, maxCol, curPos.row)
			blockerIndex := rowContains(matrix, '#', curPos.col, maxCol, curPos.row)
			if blockerIndex >= 0 && revDirIndex >= 0 && revDirIndex < blockerIndex {
				uniq++
			}
		}

		curPos = &position{row: nextPos.row, col: nextPos.col}
	}

	fmt.Println()
	matrices.DisplayMatrix(debugView)
	return uniq
}

func getRune(m [][]rune, p *position) rune {
	return m[p.row][p.col]
}
