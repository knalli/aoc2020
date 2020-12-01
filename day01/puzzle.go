package day01

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(numbers []int) error {
	matchSum := 2020
	if result := findAll2EntriesMatchingSum(numbers, matchSum); len(result) == 0 {
		return errors.New("could find entries by sum")
	} else {
		r := result[0]
		aoc.PrintSolution(fmt.Sprintf("The solution is %d * %d = %d", r.i, r.j, r.i*r.j))
	}
	return nil
}

func solve2(numbers []int) error {
	matchSum := 2020
	if result := findAll3EntriesMatchingSum(numbers, matchSum); len(result) == 0 {
		return errors.New("could find entries by sum")
	} else {
		r := result[0]
		aoc.PrintSolution(fmt.Sprintf("The solution is %d * %d * %d = %d", r.i, r.j, r.k, r.i*r.j*r.k))
	}
	return nil
}

type result2Type struct {
	i int
	j int
}

type result3Type struct {
	i int
	j int
	k int
}

func findAll2EntriesMatchingSum(numbers []int, matchSum int) []result2Type {
	result := make([]result2Type, 0)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == matchSum {
				result = append(result, result2Type{numbers[i], numbers[j]})
			}
		}
	}
	return result
}

func findAll3EntriesMatchingSum(numbers []int, matchSum int) []result3Type {
	result := make([]result3Type, 0)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == matchSum {
					result = append(result, result3Type{numbers[i], numbers[j], numbers[k]})
				}
			}
		}
	}
	return result
}
