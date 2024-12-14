package main

import (
	"fmt"
	// "math"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("Expected PROG FILENAME.")
	}

	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("Error opening file!")
	}

	fileContents := strings.Fields(string(dat))

	solutionOne := solvePartOne(fileContents)
	solutionTwo := solvePartTwo(string(dat))

	fmt.Printf("part one: %d\n", solutionOne)
	fmt.Printf("part two: %d\n", solutionTwo)
}

func solvePartOne(contents []string) *big.Int {
	maxParamVal := 999
	opName := "mul"
	opFn := func(a, b int) *big.Int {
		if a > maxParamVal || b > maxParamVal {
			return big.NewInt(0)
		}
		return big.NewInt(int64(a * b))
	}

	mulRegex := regexp.MustCompile(opName + `\(\d+,\d+\)`)
	numRegex := regexp.MustCompile(`\d+`)

	sumTotal := big.NewInt(0)
	for _, line := range contents {
		matches := mulRegex.FindAllString(line, -1)
		for _, m := range matches {
			numbers := numRegex.FindAllString(m, -1)
			a, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			sumTotal.Add(sumTotal, opFn(a, b))
		}
	}

	return sumTotal
}

func solvePartTwo(contents string) *big.Int {
	fmt.Printf("Contents: %v\n", contents)
	maxParamVal := 999
	opName := "mul"
	opFn := func(a, b int) *big.Int {
		if a > maxParamVal || b > maxParamVal {
			return big.NewInt(0)
		}
		return big.NewInt(int64(a * b))
	}

	mulRegex := regexp.MustCompile(opName + `\(\d+,\d+\)`)
	numRegex := regexp.MustCompile(`\d+`)

	deleteAfterLastDontRegex := regexp.MustCompile(`((\n|.)*)don't\(\)(\n|.)*`)
	endTrimmed := deleteAfterLastDontRegex.FindAllStringSubmatch(contents, -1)
	fmt.Printf("end trimmed: %v\n", endTrimmed)
	removeSectionsBetweenDontDoRegex := regexp.MustCompile(`(.*)don't\(\).*do\(\)(.*)`)
	allTrimmed := removeSectionsBetweenDontDoRegex.FindAllStringSubmatch(endTrimmed[0][1], -1)
	fmt.Println()
	fmt.Printf("all trimmed: %v\n", allTrimmed[0][1:])

	sumTotal := big.NewInt(0)
	for _, tuple := range allTrimmed {
		line := tuple[1]
		mulMatches := mulRegex.FindAllString(line, -1)
		for _, m := range mulMatches {
			numbers := numRegex.FindAllString(m, -1)
			a, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			sumTotal.Add(sumTotal, opFn(a, b))
		}
	}
	return sumTotal
}
