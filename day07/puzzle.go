package day07

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func solve1(lines []string) error {
	if rules, err := parseRules(lines); err != nil {
		return err
	} else {
		count := countHolders(rules, "shiny gold")
		aoc.PrintSolution(fmt.Sprintf("%d bag colors can at least contain 1 shiny gold one", count))
	}
	return nil
}

func solve2(lines []string) error {
	if rules, err := parseRules(lines); err != nil {
		return err
	} else {
		count := countBags(rules, "shiny gold") - 1
		aoc.PrintSolution(fmt.Sprintf("%d bags are required for 1 shiny gold one", count))
	}
	return nil
}

type ruleType struct {
	outerBag  string
	innerBags map[string]int
}

func (r ruleType) getInnerBagCount(bag string) int {
	if v, found := r.innerBags[bag]; found {
		return v
	} else {
		return 0
	}
}

func newRule(outerBag string, innerBags map[string]int) *ruleType {
	return &ruleType{outerBag: outerBag, innerBags: innerBags}
}

func parseRules(lines []string) ([]ruleType, error) {
	rules := make([]ruleType, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		bag := strings.Join(split[0:2], " ")
		innerBags := make(map[string]int)
		if strings.Contains(line, "contain no other bags") {
			rules = append(rules, *newRule(bag, innerBags))
			continue
		}
		for _, r := range strings.Split(strings.TrimSpace(strings.Join(split[4:], " ")), ",") {
			s := strings.Split(strings.TrimSpace(r), " ")
			innerBags[strings.Join(s[1:3], " ")] = aoc.ParseInt(s[0])
		}
		rules = append(rules, *newRule(bag, innerBags))
	}
	return rules, nil
}

func countHolders(rules []ruleType, bag string) int {
	leafs := make([]ruleType, 0)
	for _, r := range rules {
		if _, found := r.innerBags[bag]; found {
			leafs = append(leafs, r)
		}
	}

	result := make(map[string]bool)
	for len(leafs) > 0 {
		top := leafs[0]

		// truncate
		copy(leafs[0:], leafs[0+1:])
		leafs[len(leafs)-1] = ruleType{"", map[string]int{}}
		leafs = leafs[:len(leafs)-1]

		if _, found := result[top.outerBag]; found {
			continue
		}

		result[top.outerBag] = true
		for _, r := range rules {
			if _, found := r.innerBags[top.outerBag]; found {
				leafs = append(leafs, r)
			}
		}
	}
	return len(result)
}

func countBags(rules []ruleType, bag string) int {
	var rule ruleType
	for _, r := range rules {
		if r.outerBag == bag {
			rule = r
		}
	}

	result := 1
	for b, v := range rule.innerBags {
		result += v * countBags(rules, b)
	}

	return result
}
