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
		count += countDistinctGroupAnswers(answer)
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
		count += countSameGroupAnswers(answer)
	}
	aoc.PrintSolution(fmt.Sprintf("The sum of all counts is %d", count))
	return nil
}

func parseAnswers(lines []string) ([][]string, error) {
	re := regexp.MustCompile("(?:(?:\\w+\\n?)+\\s*)")
	matches := re.FindAllString(strings.Join(lines, "\n"), -1)
	result := make([][]string, len(matches))
	for i, group := range matches {
		persons := strings.Split(strings.TrimSpace(group), "\n")
		result[i] = make([]string, len(persons))
		for j, person := range persons {
			result[i][j] = person
		}
	}
	return result, nil
}

func countDistinctGroupAnswers(lines []string) int {
	counts := make(map[int32]bool)
	for _, line := range lines {
		for _, c := range line {
			if _, found := counts[c]; !found {
				counts[c] = true
			}
		}
	}
	return len(counts)
}

func countSameGroupAnswers(lines []string) int {
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
	// count only if number of counts is equal to number of answers
	r := 0
	for _, v := range counts {
		if v == len(lines) {
			r++
		}
	}
	return r
}
