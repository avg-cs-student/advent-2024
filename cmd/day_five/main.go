package main

import (
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"

	ds "github.com/avg-cs-student/advent-2024/pkg/data_structures"
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

    lines := strings.Fields(string(data))
    solutionOne := solvePartOne(lines)
    fmt.Printf("Part one: %d\n", solutionOne)
    solutionTwo := solvePartTwo(lines)
    fmt.Printf("Part two: %d\n", solutionTwo)
}

func solvePartOne(lines []string) int {
    total := 0
    rules, updates := getRulesAndUpdates(lines)

    ruleMap := map[int][]int{}
    for _, val := range rules {
        target, dependency := parseRule(val)
        ruleMap[target] = append(ruleMap[target], dependency)
    }

    for _, rawUpdate := range updates {
        seen := []int{}
        update := parseUpdate(rawUpdate)
        numOk := 0

        for i := range update {
            target := update[i]
            seen = append(seen, target)

            if deps, ok := ruleMap[target]; ok {
                depsSatisfied := true
                for _, d := range deps {
                    inUpdates := slices.Contains(update, d)
                    alreadySeen := slices.Contains(seen, d)

                    if inUpdates && !alreadySeen {
                        depsSatisfied = false
                    }
                }
                if depsSatisfied {
                    numOk++
                }
            } else {
                numOk++
            }

            if numOk == len(update) {
                center := (len(update) - 1) / 2
                total += update[center]
            }
        }
    }

    return total
}

func solvePartTwo(lines []string) int {
    total := 0
    rules, updates := getRulesAndUpdates(lines)

    ruleMap := map[int][]int{}
    for _, val := range rules {
        target, dependency := parseRule(val)
        ruleMap[target] = append(ruleMap[target], dependency)
    }

    for _, rawUpdate := range updates {
        seen := []int{}
        update := parseUpdate(rawUpdate)

        depsSatisfied := true
        outer:
        for i := range update {
            target := update[i]
            seen = append(seen, target)

            if deps, ok := ruleMap[target]; ok {
                for _, d := range deps {
                    inUpdates := slices.Contains(update, d)
                    alreadySeen := slices.Contains(seen, d)

                    if inUpdates && !alreadySeen {
                        depsSatisfied = false
                        break outer
                    }
                }
            }
        }
        if depsSatisfied {
            continue
        }

        total += reOrder(update, ruleMap)
    }

    return total
}

func reOrder(update []int, rules map[int][]int) int {
    weights := map[int]int{}
    cmp := func(a, b int) int {
        if weights[a] < weights[b] {
            return -1
        } else if weights[a] > weights[b] {
            return 1
        }

        return 0
    }
    heap := ds.NewHeap(cmp)

    for _, val := range update {
        weights[val] = 0
        if deps, ok := rules[val]; ok {
            for _, d := range deps {
                if slices.Contains(update, d) {
                    weights[val]++
                }
            }
        }
        heap.Insert(val)
    }
    
    reordered, err := heap.Dump()
    if err != nil {
        panic(err)
    } 
    center := (len(reordered) - 1) / 2
    return reordered[center]
}

func getRulesAndUpdates(lines []string) ([]string, []string) {
    var rules, updates []string
    for i := range lines {
        if strings.Contains(lines[i], ",") {
            updates = lines[i:]
            break
        }
        rules = append(rules, lines[i])
    }

    return rules, updates
}

func parseRule(rule string) (int, int) {
    components := strings.Split(rule, "|")
    dependency, err := strconv.Atoi(components[0])
    if err != nil {
        panic("Expected rules to have int|int format.")
    }

    target, err := strconv.Atoi(components[1])
    if err != nil {
        panic("Expected rules to have int|int format.")
    }

    return target, dependency
}

func parseUpdate(update string) []int {
    pages := []int{}
    for _, val := range strings.Split(update, ",") {
        num, err := strconv.Atoi(val)
        if err != nil {
            panic("Expected update to contain only ints.")
        }
        pages = append(pages, num)
    }

    return pages
}
