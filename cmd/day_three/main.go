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
	okToParse := true
	for _, line := range contents {
		mulMatches := mulRegex.FindAllIndex([]byte(line), -1)
		doMatches := doRegex.FindAllIndex([]byte(line), -1)
		dontMatches := dontRegex.FindAllIndex([]byte(line), -1)

		fmt.Printf("do(): %v\n", doMatches)
		fmt.Printf("dont(): %v\n", dontMatches)
		fmt.Printf("mul(): %v\n", mulMatches)

		nextDoStartIndex, nextDontStartIndex := math.MaxInt, math.MaxInt
		doSliceIndex, dontSliceIndex := 0, 0
		if len(doMatches) > 0 {
			nextDoStartIndex = doMatches[doSliceIndex][0]
		}

		if len(dontMatches) > 0 {
			nextDontStartIndex = dontMatches[dontSliceIndex][0]
		}

		for i := 0; i < len(mulMatches); i++ {
			mulStartIndex, mulStopIndex := mulMatches[i][0], mulMatches[i][1]

			fmt.Printf("Next do(): %d\n", nextDoStartIndex)
			fmt.Printf("Next dont(): %d\n", nextDontStartIndex)
			fmt.Printf("mul start: %d\n", mulStartIndex)

			if !okToParse && mulStartIndex < nextDoStartIndex {
				continue
			}

			if okToParse && mulStartIndex > nextDontStartIndex {
				dontSliceIndex++
				if len(dontMatches) > dontSliceIndex {
					nextDontStartIndex = dontMatches[dontSliceIndex][0]
				} else {
					nextDontStartIndex = math.MaxInt
				}
				fmt.Println("dont++")
				fmt.Printf("Next dont(): %d\n", nextDontStartIndex)
				okToParse = false
				continue
			}

			if !okToParse && mulStartIndex > nextDoStartIndex {
				doSliceIndex++
				if len(doMatches) > doSliceIndex {
					nextDoStartIndex = doMatches[doSliceIndex][0]
				} else {
					nextDoStartIndex = math.MaxInt
				}
				okToParse = true
				fmt.Println("do++")
				fmt.Printf("Next do(): %d\n", nextDoStartIndex)
			}

			if okToParse && mulStartIndex < nextDontStartIndex {
				m := line[mulStartIndex:mulStopIndex]
				fmt.Printf("Ok to parse [%d]: %s\n", mulStartIndex, m)
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
	}

	return sumTotal
}
