package day10

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"sort"
)

func solve1(lines []string) error {
	if outJolt, _, diffs, err := orderAdapters(aoc.ParseStringToIntArray(lines), 3); err != nil {
		return err
	} else {
		aoc.PrintSolution(fmt.Sprintf("out-jolt = %d", outJolt))
		aoc.PrintSolution(fmt.Sprintf("count(1-jolt-diffs) * count(3-jolt-diffs) => %d*%d = %d", diffs[1], diffs[3], diffs[1]*diffs[3]))
		return nil
	}
}

func solve2(lines []string) error {
	if _, chain, _, err := orderAdapters(aoc.ParseStringToIntArray(lines), 3); err != nil {
		return err
	} else if variants, err := countVariants(chain, 3); err != nil {
		return err
	} else {
		aoc.PrintSolution(fmt.Sprintf("variants = %d", variants))
		return nil
	}
}

func orderAdapters(adapters []int, maxDiff int) (int, []int, map[int]int, error) {
	sort.Ints(adapters)

	chain := make([]int, 0)
	diffs := make(map[int]int)
	current := 0
	chain = append(chain, 0)

	for _, n := range adapters {
		d := n - current
		current = n

		chain = append(chain, current)

		// update diff counter
		if _, found := diffs[d]; !found {
			diffs[d] = 0
		}
		diffs[d] = diffs[d] + 1
		//fmt.Printf("Put %d, d = %d\n", n, d)
	}

	// always
	diffs[maxDiff]++
	current += maxDiff
	chain = append(chain, current)

	return current, chain, diffs, nil
}

func countVariants(chain []int, maxDiff int) (int, error) {
	result := 1 // always itself

	for i := 1; i < len(chain)-1; i++ {
		// is a set of at least
		if chain[i]-chain[i-1] == 1 {
			l := 0
			for j := i + 1; j < j+4; j++ {
				if chain[j]-chain[j-1] == 1 {
					l++
				} else {
					break
				}
			}
			// l is the total of the inner elements (not the bounds)
			if l > 0 {
				// 2^l 			are all possibilities
				// 2^(l-3) 		are all possibilities which cannot be skipped
				result *= int(math.Pow(2, float64(l))) - int(math.Pow(2, float64(l-maxDiff)))
				i += l
			}
		}
	}

	return result, nil
}
