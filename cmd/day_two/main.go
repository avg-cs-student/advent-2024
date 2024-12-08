package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("Expected filename arg.")
	}
	dat, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	solutionOne := solvePartOne(lines)
	solutionTwo := solvePartTwo(lines)

	fmt.Println("Day Two")
	fmt.Printf("Part 1: %d\n", solutionOne)
	fmt.Printf("Part 2: %d\n", solutionTwo)
}

func solvePartOne(lines []string) int {
	safeLines := 0
	for _, line := range lines {
		rawVals := strings.Fields(line)
		var levels []int
		for _, val := range rawVals {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			levels = append(levels, num)
		}

		var cmp func(a, b int) bool
		if levels[0] - levels[1] > 0 {
			cmp = func(a, b int) bool {
				diff := a - b
				if diff < 0 {
					diff = 0 - diff
				}
				return a > b && diff < 4
			}
		} else {
			cmp = func(a, b int) bool {
				diff := a - b
				if diff < 0 {
					diff = 0 - diff
				}
				return a < b && diff < 4
			}
		}
		if isSafe(levels, cmp, 0) {
			safeLines++
		}
	}
	return safeLines
}

func isSafe(levels []int, cmp func(a, b int) bool, tolerance int) bool {
	var check func([]int, int) bool
	check = func(levels []int, bad int) bool {
		for i := 1; i < len(levels); i++ {
			if cmp(levels[i - 1], levels[i]) {
				continue
			}

			bad++
			if bad > tolerance {
				return false
			}

			if i == 1 {
				continue
			}

			if i < len(levels) - 1 {
				if cmp(levels[i - 1], levels[i + 1]) {
					i++
					continue
				}
				if i > 1 {
					if cmp(levels[i - 2], levels[i + 1]) {
						i++
						continue
					}
				}
			} else {
				continue
				// return true
			}

			return false
		}
		return true
	}

	return check(levels, 0)
}

func solvePartTwo(lines []string) int {
	safeLines := 0
	for _, line := range lines {
		rawVals := strings.Fields(line)
		var levels []int
		for _, val := range rawVals {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			levels = append(levels, num)
		}

		cmp1 := func(a, b int) bool {
			diff := a - b
			if diff < 0 {
				diff = 0 - diff
			}

			return a > b && diff < 4
		}
		cmp2 := func(a, b int) bool {
			diff := a - b
			if diff < 0 {
				diff = 0 - diff
			}

			return a < b && diff < 4
		}

		// just check both, its cheap :)
		if isSafe(levels, cmp1, 1) || isSafe(levels, cmp2, 1) {
			safeLines++
		}
	}
	return safeLines
}
