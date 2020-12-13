package day13

import "github.com/knalli/aoc"

func MustParseIntsIgnore(lines []string, filter func(s string) bool) []int {
	result := make([]int, 0)
	for _, line := range lines {
		if filter(line) {
			result = append(result, aoc.ParseInt(line))
		}
	}
	return result
}

func MustParseIntsIgnoreWithAlternative(lines []string, alternate int, filter func(s string) bool) []int {
	result := make([]int, 0)
	for _, line := range lines {
		if filter(line) {
			result = append(result, aoc.ParseInt(line))
		} else {
			result = append(result, alternate)
		}
	}
	return result
}
