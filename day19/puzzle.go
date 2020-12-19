package day19

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func solve1(lines []string) error {
	rules, messages := splitInput(lines)
	count := 0
	for _, message := range messages {
		if isValid(message, rules) {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Number of messages matching 'rule 0' is %d", count))
	return nil
}

func solve2(lines []string) error {
	rules, messages := splitInput(lines)
	rules[8].raw = "42 | 42 8"
	rules[8].groups = parseMultiGroup("42 | 42 8")
	rules[11].raw = "42 31 | 42 11 31"
	rules[11].groups = parseMultiGroup("42 31 | 42 11 31")
	count := 0
	for _, message := range messages {
		if isValid(message, rules) {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Number of messages matching 'rule 0' is %d", count))
	return nil
}

type Rule struct {
	id         int
	raw        string
	groups     [][]int
	match      bool
	matchValue uint8
}

func NewRule(raw string) Rule {
	r := Rule{raw: raw}
	if raw[0] == '"' {
		r.match = true
		r.matchValue = raw[1]
	} else {
		r.groups = parseMultiGroup(raw)
	}
	return r
}

func parseMultiGroup(value string) [][]int {
	groups := make([][]int, 0)
	for _, group := range strings.Split(strings.TrimSpace(value), "|") {
		groups = append(groups, parseGroup(group))
	}
	return groups
}

func parseGroup(group string) []int {
	return aoc.ParseInts(strings.TrimSpace(group), " ")
}

func splitInput(lines []string) (rules map[int]*Rule, messages []string) {
	rules = make(map[int]*Rule)
	messages = make([]string, 0)
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.Contains(line, ":") {
			idx := strings.Index(line, ": ")
			ruleId := aoc.ParseInt(line[0:idx])
			rule := NewRule(line[idx+2:])
			rule.id = ruleId
			rules[ruleId] = &rule
		} else {
			messages = append(messages, line)
		}
	}
	return
}

func isValid(message string, rules map[int]*Rule) bool {

	var valid func(pos int, rule *Rule) ([]int, bool)
	valid = func(pos int, rule *Rule) ([]int, bool) {
		if pos > len(message)-1 {
			return []int{}, false
		}
		if rule.match {
			ok := message[pos] == rule.matchValue
			if ok {
				//printMessagePosition(message, pos)
			}
			return []int{pos + 1}, ok
		}
		lastPositions := make([]int, 0)
		for _, group := range rule.groups {
			groupStartPos := pos
			groupRunningCurrentPositions := []int{groupStartPos}
			for _, subRuleId := range group {
				subRule := rules[subRuleId]
				groupRunningNextPositions := make([]int, 0)
				for _, groupRunningPos := range groupRunningCurrentPositions {
					if subGroupLastPositions, subGroupValid := valid(groupRunningPos, subRule); subGroupValid {
						for _, subGroupLastPos := range subGroupLastPositions {
							groupRunningNextPositions = append(groupRunningNextPositions, subGroupLastPos)
						}
					} else {
						continue
					}
				}
				if len(groupRunningNextPositions) == 0 {
					groupRunningCurrentPositions = []int{}
					break
				}
				groupRunningCurrentPositions = groupRunningNextPositions
			}
			for _, p := range groupRunningCurrentPositions {
				lastPositions = append(lastPositions, p)
			}
		}
		if len(lastPositions) > 0 {
			return lastPositions, true
		}
		return []int{pos}, false
	}

	if lastPositions, ok := valid(0, rules[0]); ok && len(lastPositions) > 0 {
		for _, p := range lastPositions {
			if p == len(message) {
				return true
			}
		}
	}
	return false
}
