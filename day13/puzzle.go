package day13

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

func solve1(lines []string) error {
	initialWaitTime := aoc.ParseInt(lines[0])
	busLines := MustParseIntsIgnore(strings.Split(lines[1], ","), func(s string) bool {
		return s != "x"
	})
	if busId, waitTime, err := resolveFirstMatch(initialWaitTime, busLines); err != nil {
		return err
	} else {
		aoc.PrintSolution(fmt.Sprintf("The earliest bus #%d has a waiting time of %dm: %d", busId, waitTime, busId*waitTime))
	}
	return nil
}

func solve2(lines []string) error {
	busLines := MustParseIntsIgnoreWithAlternative(strings.Split(lines[1], ","), -1, func(s string) bool {
		return s != "x"
	})
	result := resolveGoldCoinCompetition(0, busLines)
	aoc.PrintSolution(fmt.Sprintf("Earliest matching timestamp = %d", result))
	return nil
}

func resolveFirstMatch(offset int, ids []int) (resultId int, waitTime int, err error) {
	waitTime = math.MaxInt64
	resultId = -1
	for _, id := range ids {
		wait := id - (offset % id)
		if wait < waitTime {
			waitTime = wait
			resultId = id
		}
	}
	return
}

func resolveGoldCoinCompetition(offset int, ids []int) int {
	t := offset
	delta := 1
	txm := -1

	for {
		valid := true
		for idx, id := range ids {
			if !(id > 0) {
				continue // next id
			}
			i := (t + (idx)) % (id)
			if i != 0 {
				valid = false
				break // next round(t)
			}
			if idx > txm {
				//fmt.Printf("[%20d] [%20d]   +%2d   #%d\n", t, t+(idx), idx, id)
				txm = idx
				// This is the actual gem: ensures all next steps are a multiple so skipping most
				delta *= id
			}
		}
		if valid {
			return t
		}
		t += delta
	}
}
