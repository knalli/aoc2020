package day16

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func solve1(lines []string) error {
	rules, _, nearby := parseInput(lines)
	sum := 0
	for _, t := range nearby {
		isTicketValid(t, rules, func(v int) {
			sum += v
		})
	}
	aoc.PrintSolution(fmt.Sprintf("Ticket scanning error rate: %d", sum))
	return nil
}

func solve2(lines []string) error {
	rules, myTicket, nearby := parseInput(lines)
	validTickets := make([][]int, 0)
	for _, t := range nearby {
		if isTicketValid(t, rules, nil) {
			validTickets = append(validTickets, t)
		}
	}
	mappings := resolveTicketFields(validTickets, rules)
	//aoc.PrintSolution(fmt.Sprintf("Found %d field mappings at all", len(mappings)))
	result := 1
	for name, idx := range mappings {
		if strings.Index(name, "departure") == 0 {
			result *= myTicket[idx]
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Result of the multiplication of all departure fields is %d", result))
	return nil
}

type TicketRule struct {
	name   string
	ranges []IntRange
}

type IntRange struct {
	start int
	end   int
}

func parseInput(lines []string) ([]TicketRule, []int, [][]int) {
	var ticketRules []TicketRule
	var myTicket []int
	var nearbyTickets [][]int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "your ticket:" {
			ticketRules = parseTicketRules(lines[0 : i-1])
			continue
		}
		if lines[i] == "nearby tickets:" {
			myTicket = aoc.ParseInts(lines[i-2], ",")
			nearbyTickets = parseTicketValues(lines[i+1:])
			break
		}
	}
	return ticketRules, myTicket, nearbyTickets
}

func parseTicketRules(lines []string) []TicketRule {
	result := make([]TicketRule, len(lines))
	for i, line := range lines {
		ranges := make([]IntRange, 0)
		for _, r := range strings.Split(line[strings.Index(line, ":")+2:], " or ") {
			s := strings.Split(r, "-")
			ranges = append(ranges, IntRange{start: aoc.ParseInt(s[0]), end: aoc.ParseInt(s[1])})
		}
		result[i] = TicketRule{
			name:   line[0:strings.Index(line, ":")],
			ranges: ranges,
		}
	}
	return result
}

func parseTicketValues(lines []string) [][]int {
	result := make([][]int, len(lines))
	for i, line := range lines {
		result[i] = aoc.ParseInts(line, ",")
	}
	return result
}

func isTicketValid(values []int, rules []TicketRule, onInvalidValue func(v int)) bool {
	result := true
	for _, v := range values {
		valid := false
		for _, rule := range rules {
			for _, r := range rule.ranges {
				if r.start <= v && v <= r.end {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if !valid {
			if onInvalidValue != nil {
				onInvalidValue(v)
			}
			result = false
		}
	}
	return result
}

func resolveTicketFields(tickets [][]int, rules []TicketRule) map[string]int {
	fieldNames := make(map[string]int)
	fieldPositions := make(map[int]string)

	for len(fieldNames) < len(rules) {
		for _, rule := range rules {
			if _, exist := fieldNames[rule.name]; exist {
				continue // already mapped
			}

			positionCandidates := make([]int, 0)
			// each possible index for the field
			for p := 0; p < len(rules); p++ {
				if _, exist := fieldPositions[p]; exist {
					continue // already mapped
				}

				ticketsValid := true
				for _, ticket := range tickets {
					valid := false
					for _, r := range rule.ranges {
						if r.start <= ticket[p] && ticket[p] <= r.end {
							valid = true
						}
					}
					if !valid {
						ticketsValid = false
						break
					}
				}
				if ticketsValid {
					// candidate found
					positionCandidates = append(positionCandidates, p)
				}
			}

			// only use if exact one found
			if len(positionCandidates) == 1 {
				p := positionCandidates[0]
				fieldNames[rule.name] = p
				fieldPositions[p] = rule.name
			}

		}
	}

	return fieldNames
}
