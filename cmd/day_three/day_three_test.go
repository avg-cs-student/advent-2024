package main

import (
	"math/big"
	"testing"
)

func TestSolutionTwoStock(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := big.NewInt(48)
	if sum := solvePartTwo([]string{input}); sum.Cmp(expected) != 0 {
		t.Errorf("Expected %v got %v.\n", expected, sum)
	}
}

func TestSolutionTwoCustom1(t *testing.T) {
	input := `do()xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := big.NewInt(48)
	if sum := solvePartTwo([]string{input}); sum.Cmp(expected) != 0 {
		t.Errorf("Expected %v got %v.\n", expected, sum)
	}
}

func TestSolutionTwoCustom2(t *testing.T) {
	input := `don't()xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := big.NewInt(40)
	if sum := solvePartTwo([]string{input}); sum.Cmp(expected) != 0 {
		t.Errorf("Expected %v got %v.\n", expected, sum)
	}
}
