package day09

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"math"
)

func solve1(lines []string) error {
	if pos, n, err := decode(aoc.ParseStringToIntArray(lines), 25); err == nil {
		return errors.New("no weakness found")
	} else if err.Error() == "weakness found" {
		aoc.PrintSolution(fmt.Sprintf("Weakness found with numer %d at #%d", n, pos))
		return nil
	} else {
		return err
	}
}

func solve2(lines []string) error {
	_, weakness, _ := decode(aoc.ParseStringToIntArray(lines), 25)
	if set := findContiguousSet(aoc.ParseStringToIntArray(lines), weakness); set == nil {
		return errors.New("no weakness found")
	} else {
		min := minIntArrayValue(set)
		max := aoc.MaxIntArrayValue(set)
		aoc.PrintSolution(fmt.Sprintf("Weakness found with set for %d: %d + %d = %d", len(set), min, max, min+max))
		return nil
	}
}

func decode(input []int, window int) (int, int, error) {

	var pos int
	var scope []int
	var current int

	// initial
	pos = window

	for pos < len(input) {
		scope = input[pos-window : pos]
		current = input[pos]

		// sum?
		found := false
		for i, a := range scope {
			for j, b := range scope {
				if i == j {
					continue
				}
				if a+b == current {
					found = true
				}
			}
		}
		if !found {
			return pos, current, errors.New("weakness found")
		}

		pos++
	}

	return -1, 0, nil
}

func findContiguousSet(input []int, sum int) []int {
	for i := 0; i < len(input); i++ {
		tsum := input[i]
		for j := i + 1; j < len(input); j++ {
			tsum += input[j]
			if tsum == sum {
				return input[i : j+1]
			}
			if tsum > sum {
				break
			}
		}
	}
	return nil
}

func minIntArrayValue(arr []int) int {
	min := math.MaxInt64
	for _, n := range arr {
		min = aoc.MinInt(min, n)
	}
	return min
}
