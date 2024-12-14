package main

import (
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
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
