package day06

import (
	"fmt"
	"github.com/knalli/aoc"
	"regexp"
	"strings"
)

func solve1(lines []string) error {
	answers, err := parseAnswers(lines)
	if err != nil {
		return err
	}
	count := 0
	for _, answer := range answers {
		count += countGroupAnswers(answer, func(total int) bool {
			return true
		})
	}
	aoc.PrintSolution(fmt.Sprintf("The sum of all counts is %d", count))
	return nil
}

func solve2(lines []string) error {
	answers, err := parseAnswers(lines)
	if err != nil {
		return err
	}
	count := 0
	for _, answer := range answers {
		l := len(answer)
		count += countGroupAnswers(answer, func(total int) bool {
			return total == l
		})
	}
	aoc.PrintSolution(fmt.Sprintf("The sum of all counts is %d", count))
	return nil
}

func parseAnswers(lines []string) ([][]string, error) {
	re := regexp.MustCompile("(?:(?:\\w+\\n?)+\\s*)")
	matches := re.FindAllString(strings.Join(lines, "\n"), -1)
	result := make([][]string, len(matches))
	for i, group := range matches {
		people := strings.Split(strings.TrimSpace(group), "\n")
		result[i] = make([]string, len(people))
		for j, person := range people {
			result[i][j] = person
		}
	}
	return result, nil
}

func countGroupAnswers(lines []string, check func(total int) bool) int {
	counts := make(map[int32]int)
	for _, line := range lines {
		for _, c := range line {
			if v, found := counts[c]; found {
				counts[c] = v + 1
			} else {
				counts[c] = 1
			}
		}
	}
	r := 0
	for _, v := range counts {
		if check(v) {
			r++
		}
	}
	return r
}
