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
	//
	solutionTwo := solvePartTwo(matrix)
	fmt.Printf("Part one: %d\n", solutionTwo)
}

type Pos struct {
	x int
	y int
}

func (p Pos) calcAntinodes(other Pos) (Pos, Pos) {
	deltaX := p.x - other.x
	deltaY := p.y - other.y
	nodeOne := Pos{x: p.x + deltaX, y: p.y + deltaY}
	nodeTwo := Pos{x: other.x - deltaX, y: other.y - deltaY}
	return nodeOne, nodeTwo
}

func (p Pos) isValid(xMax, yMax int) bool {
	return p.x < xMax && p.x >= 0 && p.y < yMax && p.y >= 0
}

func (p Pos) calcExtendedAntinodes(other Pos, xMax, yMax int) []Pos {
	nodes := []Pos{}
	deltaX := p.x - other.x
	deltaY := p.y - other.y

	next := Pos{x: p.x + deltaX, y: p.y + deltaY}
	for next.isValid(xMax, yMax) {
		nodes = append(nodes, next)
		next = Pos{x: next.x + deltaX, y: next.y + deltaY}
	}
	next = Pos{x: other.x - deltaX, y: other.y - deltaY}
	for next.isValid(xMax, yMax) {
		nodes = append(nodes, next)
		next = Pos{x: next.x - deltaX, y: next.y - deltaY}
	}

	return nodes
}

func solvePartOne(matrix [][]rune) int {
	locations := map[rune][]Pos{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			currentRune := matrix[row][col]
			if currentRune == '.' {
				continue
			}
			_, found := locations[currentRune]
			if !found {
				locations[currentRune] = []Pos{
					{x: row, y: col},
				}
				continue
			}

			locations[currentRune] = append(locations[currentRune], Pos{x: row, y: col})
		}
	}
	antinodes := map[Pos]bool{}

	for _, positionList := range locations {
		for i, antennaLoc := range positionList {
			for j := i + 1; j < len(positionList); j++ {
				nodeOne, nodeTwo := antennaLoc.calcAntinodes(positionList[j])
				if nodeOne.isValid(len(matrix), len(matrix[0])) {
					antinodes[nodeOne] = true
					matrix[nodeOne.x][nodeOne.y] = '#'
				}
				if nodeTwo.isValid(len(matrix), len(matrix[0])) {
					antinodes[nodeTwo] = true
					matrix[nodeTwo.x][nodeTwo.y] = '#'
				}
			}
		}
	}
	return len(antinodes)
}

func solvePartTwo(matrix [][]rune) int {
	locations := map[rune][]Pos{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			currentRune := matrix[row][col]
			if currentRune == '.' {
				continue
			}
			_, found := locations[currentRune]
			if !found {
				locations[currentRune] = []Pos{
					{x: row, y: col},
				}
				continue
			}

			locations[currentRune] = append(locations[currentRune], Pos{x: row, y: col})
		}
	}

	antinodes := map[Pos]bool{}
	for _, positionList := range locations {
		for i, antennaLoc := range positionList {
			antinodes[antennaLoc] = true
			for j := i + 1; j < len(positionList); j++ {
				nodes := antennaLoc.calcExtendedAntinodes(positionList[j], len(matrix), len(matrix[0]))
				for _, n := range nodes {
					antinodes[n] = true
					matrix[n.x][n.y] = '#'
				}
			}
		}
	}
	return len(antinodes)
}
