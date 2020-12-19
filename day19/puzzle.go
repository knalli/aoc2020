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
		// out of bound
		if pos > len(message)-1 {
			return []int{}, false
		}
		// direct match
		if rule.match {
			ok := message[pos] == rule.matchValue
			return []int{pos + 1}, ok
		}
		// sub rules / groups matching
		lastPositions := make([]int, 0)
		for _, group := range rule.groups {
			// all position options for this group (initially the current position)
			currentPositions := []int{pos}
			for _, subRuleId := range group {
				subRule := rules[subRuleId]
				// each position will be consumed; if verified, its successor will be added again
				successorPositions := make([]int, 0)
				for _, groupRunningPos := range currentPositions {
					// only if the sub rule is valid, all its last positions will be added as successors
					if subRuleLastPositions, subRuleValid := valid(groupRunningPos, subRule); subRuleValid {
						for _, p := range subRuleLastPositions {
							successorPositions = append(successorPositions, p)
						}
					} else {
						continue // running position invalid, try next
					}
				}
				if len(successorPositions) == 0 {
					// if no successor was found, the current positions are all invalid
					currentPositions = []int{}
					break // rule invalid, try next
				} else {
					// reset the current positions with the new successor ones
					currentPositions = successorPositions
				}
			}
			// collect all group's positions as available lastPositions for this rule
			for _, p := range currentPositions {
				lastPositions = append(lastPositions, p)
			}
		}
		if len(lastPositions) > 0 {
			return lastPositions, true
		}
		return []int{pos}, false
	}

	if lastPositions, ok := valid(0, rules[0]); ok && len(lastPositions) > 0 {
		// only a lastPosition at the position right after the last valid one (i.e. len(messages)) is valid
		for _, p := range lastPositions {
			if p == len(message) {
				return true
			}
		}
	}
	return false
}
