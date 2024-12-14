package main

import (
    "fmt"
    "os"
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
}

func solvePartOne(rawContents []byte) int {
    matrix := createRuneMatrix(rawContents)
    displayMatrix(matrix)
    return check(matrix)
}

func createRuneMatrix(data []byte) [][]rune {
    var matrix [][]rune
    matrix = append(matrix, []rune{})

    row := 0
    for i, b := range data {
        switch b {
        case '\n':
            row++
            if i == len(data) - 1 {
                break
            }
            matrix = append(matrix, []rune{})
        default:
            matrix[row] = append(matrix[row], rune(b))
        }
    }

    return matrix
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
        return traverse(nextRow, nextCol, needed + 1, nextIndex)
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

func displayMatrix(matrix [][]rune) {
    for i := range matrix {
        for j := range matrix[i] {
            fmt.Printf("%c ", matrix[i][j])
        }
        fmt.Println()
    }
}
