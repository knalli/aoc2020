package day02

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type policyType struct {
	str string
	min int
	max int
}

type passwordValidator func(password string, policy *policyType) bool

func solve1(lines []string) error {
	count := 0
	for _, line := range lines {
		if valid, err := isPasswordValid(line, step1Validator); err != nil {
			return err
		} else if valid {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("%d of %d passwords are valid", count, len(lines)))
	return nil
}

func solve2(lines []string) error {
	count := 0
	for _, line := range lines {
		if valid, err := isPasswordValid(line, step2Validator); err != nil {
			return err
		} else if valid {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("%d of %d passwords are valid", count, len(lines)))
	return nil
}

func isPasswordValid(password string, validator passwordValidator) (bool, error) {
	split := strings.Split(password, ":")
	password = strings.TrimSpace(split[1])
	policyStr := strings.TrimSpace(split[0])

	if policy, err := parsePolicy(policyStr); err != nil {
		return false, err
	} else {
		return validator(password, policy), nil
	}
}

func step1Validator(password string, policy *policyType) bool {
	found := 0
	for _, r := range password {
		c := string(r)
		if c == policy.str {
			found++
		}
	}
	return policy.min <= found && found <= policy.max
}

func step2Validator(password string, policy *policyType) bool {
	a := string(password[policy.min-1]) == policy.str
	b := string(password[policy.max-1]) == policy.str
	return (a && !b) || (!a && b)
}

func parsePolicy(str string) (*policyType, error) {
	p1 := strings.Split(str, " ")
	if len(p1) != 2 {
		return nil, errors.New("invalid policy")
	}
	p2 := strings.Split(p1[0], "-")
	if len(p2) != 2 {
		return nil, errors.New("invalid policy")
	}
	return &policyType{str: p1[1], min: aoc.ParseInt(p2[0]), max: aoc.ParseInt(p2[1])}, nil
}
