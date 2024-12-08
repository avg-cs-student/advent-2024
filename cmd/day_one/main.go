package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	ds "github.com/avg-cs-student/advent-2024/pkg/data_structures"
)

func main() {
	if len(os.Args) != 2 {
		panic("Not enough args!!")
	}
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("Error opening file!")
	}

	fileContents := strings.Fields(string(dat))

	solutionOne := solvePartOne(fileContents)
	fmt.Printf("Part one: %d\n", solutionOne)

	solutionTwo := solvePartTwo(fileContents)
	fmt.Printf("Part two: %d\n", solutionTwo)
}

func solvePartOne(contents []string) int64 {
	minHeap := func(a, b int) int {
		if a < b {
			return 1
		}
		if a > b {
			return -1
		}

		return 0
	}

	listOne := ds.NewHeap(minHeap)
	listTwo := ds.NewHeap(minHeap)
	for i, num := range contents {
		if i%2 == 0 {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			listOne.Insert(val)
		} else {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			listTwo.Insert(val)
		}
	}

	var sumTotal int64
	a, _ := listOne.Dump()
	b, _ := listTwo.Dump()
	for i := range b {
		rawDiff := a[i] - b[i]
		if rawDiff < 0 {
			rawDiff = 0 - rawDiff
		}
		sumTotal += int64(rawDiff)
	}

	return sumTotal
}

func solvePartTwo(contents []string) int64 {
	listOne := []int{}
	listTwoMap := map[int]int{}

	for i, num := range contents {
		if i%2 == 0 {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			listOne = append(listOne, val)
		} else {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			listTwoMap[val]++
		}
	}

	var sumTotal int64
	for _, val := range listOne {
		similarity := val * listTwoMap[val]
		sumTotal += int64(similarity)
	}

	return sumTotal
}
