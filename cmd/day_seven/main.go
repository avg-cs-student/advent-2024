package main

import (
	"bufio"
	"fmt"
	"os"
    "slices"
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
	fmt.Printf("Part 1: %d\n", solutionOne)

	solutionTwo := solvePartTwo(lines)
	fmt.Printf("Part 2: %d\n", solutionTwo)
}

var (
    validOps = []func(a, b int) int {
        func(a, b int) int { return a * b },
        func(a, b int) int { return a + b },
    }
)

func solvePartOne(lines []string) int {
    total := 0
    for _, line := range lines {
        raw := strings.Split(line, ":")
        expected, err := strconv.Atoi(raw[0])
        if err != nil {
            panic(err)
        }

        vals := []int{}
        for _, raw := range strings.Fields(raw[1]) {
            num, err := strconv.Atoi(raw)
            if err != nil {
                panic(err)
            }
            vals = append(vals, num)
        }

        ok := solution(expected, vals)
        if ok {
            total += expected
        }
    }
    return total
}

const (
    add = iota
    mult
)

func solution(expected int, vals []int) bool {
    cache := [][]int{}
    
    for i := 0; i < len(vals); i++ {
        cache = append(cache, []int{})
        if i == 0 {
            cache[i] = append(cache[i], vals[i])
            continue
        }
        for j := 0; j < len(cache[i - 1]); j++ {
            cache[i] = append(cache[i], cache[i - 1][j] + vals[i], cache[i - 1][j] * vals[i])
        }
    }
    return slices.Contains(cache[len(vals) - 1], expected)
}

func solvePartTwo(lines []string) int {
    total := 0
    for _, line := range lines {
        raw := strings.Split(line, ":")
        expected, err := strconv.Atoi(raw[0])
        if err != nil {
            panic(err)
        }

        vals := []int{}
        for _, raw := range strings.Fields(raw[1]) {
            num, err := strconv.Atoi(raw)
            if err != nil {
                panic(err)
            }
            vals = append(vals, num)
        }

        ok := line_solution_2(expected, vals)
        if ok {
            total += expected
        }
    }
    return total
}

func line_solution_2(expected int, vals []int) bool {
    cache := [][]int{}
    
    for i := 0; i < len(vals); i++ {
        cache = append(cache, []int{})
        if i == 0 {
            cache[i] = append(cache[i], vals[i])
            continue
        }
        for j := 0; j < len(cache[i - 1]); j++ {
            cache[i] = append(cache[i], cache[i - 1][j] + vals[i], cache[i - 1][j] * vals[i], concat(cache[i - 1][j], vals[i]))
        }
    }
    return slices.Contains(cache[len(vals) - 1], expected)
}

func concat(a, b int) int {
    aStr := strconv.Itoa(a)
    bStr := strconv.Itoa(b)

    c, err := strconv.Atoi(aStr + bStr)
    if err != nil {
        panic(err)
    } 
    return c
}
