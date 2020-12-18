package day18

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	sum := 0
	for _, line := range lines {
		//fmt.Printf("Expr: %s", line)
		if result, err := evaluate(parse(line), false); err != nil {
			//fmt.Printf(" << Error: %s\n", err.Error())
		} else {
			sum += result
			//fmt.Printf(" = %d\n", result)
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Sum of all: %d", sum))
	return nil
}

func solve2(lines []string) error {
	sum := 0
	for _, line := range lines {
		//fmt.Printf("Expr: %s", line)
		if result, err := evaluate(parse(line), true); err != nil {
			//fmt.Printf(" << Error: %s\n", err.Error())
		} else {
			sum += result
			//fmt.Printf(" = %d\n", result)
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Sum of all: %d", sum))
	return nil
}

const OPERATOR = "OPERATOR"
const NUMBER = "NUMBER"
const BRACE_OPEN = "BRACE_OPEN"
const BRACE_CLOSE = "BRACE_CLOSE"

const OPERATION_ADD = 0
const OPERATION_MULTI = 1

type token struct {
	Type  string
	Value int
}

func parse(expr string) []token {

	result := make([]token, 0)

	for _, c := range expr {
		if c == ' ' {
			continue
		} else if c == '(' {
			result = append(result, token{Type: BRACE_OPEN})
		} else if c == ')' {
			result = append(result, token{Type: BRACE_CLOSE})
		} else if c == '*' {
			result = append(result, token{Type: OPERATOR, Value: OPERATION_MULTI})
		} else if c == '+' {
			result = append(result, token{Type: OPERATOR, Value: OPERATION_ADD})
		} else if '0' <= c && c <= '9' {
			n := int(c - 48)
			result = append(result, token{Type: NUMBER, Value: n})
		}
	}

	return result
}

func evaluate(tokens []token, preferAddOperation bool) (int, error) {

	var eval func([]token) token
	eval = func(exp []token) (result token) {
		result.Type = NUMBER

		// easy?
		if len(exp) == 1 {
			result.Value = exp[0].Value
			return
		}

		// solve plus operations at first (but skip too simple expressions)
		if preferAddOperation && len(exp) > 3 {
			for {
				opAddIdx := -1
				for i, t := range exp {
					if t.Type == OPERATOR && t.Value == OPERATION_ADD {
						opAddIdx = i
						break
					}
				}
				if opAddIdx <= -1 {
					// no more add-operation
					break
				}
				// solve sub expression "A + B"
				start := opAddIdx - 1
				next := make([]token, len(exp)-2)
				// copy everything before the sub expression
				for i := 0; i < start; i++ {
					next[i] = exp[i]
				}
				// evaluate sub expression
				next[start] = eval(exp[start : start+3])
				// copy everything after the sub expression
				for i := start + 3; i < len(exp); i++ {
					next[i-2] = exp[i]
				}
				exp = next
			}
		}

		// easy?
		if len(exp) == 1 {
			result.Value = exp[0].Value
			return
		}

		lastOp := 0
		for i, t := range exp {
			if t.Type == OPERATOR {
				lastOp = t.Value
			} else if t.Type == NUMBER {
				if i == 0 {
					// initial
					result.Value = t.Value
				} else {
					// consecutive
					if lastOp == OPERATION_ADD {
						result.Value += t.Value
					} else if lastOp == 1 {
						result.Value *= t.Value
					}
				}
			}
		}
		return
	}

	// solve iteratively all sub expressions (form inner to outer)
	for {
		lastOpenIdx := -1

		for i, t := range tokens {
			if t.Type == BRACE_OPEN {
				// overwrite any found open (maybe next time)
				lastOpenIdx = i
			} else if t.Type == BRACE_CLOSE {
				// at this point, there was found the most inner block
				next := make([]token, 0)
				// copy everything before the open brace
				for j := 0; j < lastOpenIdx; j++ {
					next = append(next, tokens[j])
				}
				// evaluate the sub expression (does not include any braces)
				next = append(next, eval(tokens[lastOpenIdx+1:i]))
				// copy everything after the closing brace
				for j := i + 1; j < len(tokens); j++ {
					next = append(next, tokens[j])
				}
				// replace, try agian
				tokens = next
				break
			}
		}

		if lastOpenIdx == -1 {
			break
		}
	}

	return eval(tokens).Value, nil
}
