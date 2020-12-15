package day15

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	maxTurn := 2020
	result := playGameUntilTurn(aoc.ParseInts(lines[0], ","), maxTurn)
	aoc.PrintSolution(fmt.Sprintf("The %dth number is %d", maxTurn, result))
	return nil
}

func solve2(lines []string) error {
	maxTurn := 30000000
	result := playGameUntilTurn(aoc.ParseInts(lines[0], ","), maxTurn)
	aoc.PrintSolution(fmt.Sprintf("The %dth number is %d", maxTurn, result))
	return nil
}

type delta struct {
	diff int
	last int
}

func playGameUntilTurn(initial []int, maxTurn int) int {
	spokens := make(map[int]int)
	deltas := make(map[int]delta)

	turn := 1
	for _, n := range initial {
		spokens[turn] = n
		deltas[n] = delta{diff: 0, last: turn}
		turn++
	}
	for {
		lastSpoken := spokens[turn-1]
		speak := deltas[lastSpoken].diff
		spokens[turn] = speak
		if d, found := deltas[speak]; found {
			deltas[speak] = delta{diff: turn - d.last, last: turn}
		} else {
			deltas[speak] = delta{diff: 0, last: turn}
		}
		if turn == maxTurn {
			return spokens[maxTurn]
		}
		turn++
	}
}
