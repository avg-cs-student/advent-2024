package main

import (
	"fmt"
	"math"
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
	solutionTwo := solvePartTwo(fileContents)

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

func solvePartTwo(contents []string) *big.Int {
	maxParamVal := 999
	opName := "mul"
	opFn := func(a, b int) *big.Int {
		if a > maxParamVal || b > maxParamVal {
			return big.NewInt(0)
		}
		return big.NewInt(int64(a * b))
	}

	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)
	mulRegex := regexp.MustCompile(opName + `\(\d+,\d+\)`)
	numRegex := regexp.MustCompile(`\d+`)

	sumTotal := big.NewInt(0)
	okayUntil := math.MaxInt
	okayAfter := 0
	for _, line := range contents {
		mulMatches := mulRegex.FindAllIndex([]byte(line), -1)
		doMatches := doRegex.FindAllIndex([]byte(line), -1)
		dontMatches := dontRegex.FindAllIndex([]byte(line), -1)

		if len(dontMatches) > 0 {
			okayUntil = dontMatches[0][0]
		}
		if len(doMatches) > 0 {
			okayAfter = doMatches[0][0]
		}

		for mul, do, dont := 0, 0, 0; mul < len(mulMatches); mul++ {
			start, stop := mulMatches[mul][0], mulMatches[mul][1]
			if start > okayUntil && start < okayAfter {
				continue
			} 

			if start < okayUntil && okayUntil > okayAfter {
				do++
				if len(doMatches) > do {
					okayAfter = doMatches[do][0]
				} else {
					okayAfter = math.MaxInt
				}
			} else if start > okayAfter && okayAfter > okayUntil {
				dont++
				if len(dontMatches) > dont {
					okayUntil = dontMatches[dont][0]
				} else {
					okayUntil = math.MaxInt
				}
			}

			m := line[start:stop]
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
