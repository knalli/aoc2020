package day03

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	area := parseArea(lines)
	way := runSlope(area, func(c coordinates) coordinates {
		return coordinates{x: c.x + 3, y: c.y + 1}
	})
	if count, err := countFilled(way, area); err != nil {
		return err
	} else {
		//fmt.Println(area.toString(false))
		aoc.PrintSolution(fmt.Sprintf("Tobogann will encounter %d trees", count))
	}
	return nil
}

func solve2(lines []string) error {
	area := parseArea(lines)

	navigators := make([]navigator, 5)
	navigators[0] = func(c coordinates) coordinates {
		return coordinates{x: c.x + 1, y: c.y + 1}
	}
	navigators[1] = func(c coordinates) coordinates {
		return coordinates{x: c.x + 3, y: c.y + 1}
	}
	navigators[2] = func(c coordinates) coordinates {
		return coordinates{x: c.x + 5, y: c.y + 1}
	}
	navigators[3] = func(c coordinates) coordinates {
		return coordinates{x: c.x + 7, y: c.y + 1}
	}
	navigators[4] = func(c coordinates) coordinates {
		return coordinates{x: c.x + 1, y: c.y + 2}
	}
	results := make([]int, len(navigators))
	for i, nav := range navigators {
		way := runSlope(area, nav)
		if count, err := countFilled(way, area); err != nil {
			return err
		} else {
			results[i] = count
		}
	}
	result := 1
	for _, n := range results {
		result *= n
	}
	aoc.PrintSolution(fmt.Sprintf("Tobogann will encounter %d trees", result))
	return nil
}

func countFilled(way []coordinates, area *area) (int, error) {
	count := 0
	for _, c := range way {
		if c, err := area.get(c); err != nil {
			return -1, err
		} else if c == FILLED {
			count++
		}
	}
	return count, nil
}

type navigator func(c coordinates) coordinates

func runSlope(a *area, nav navigator) []coordinates {

	result := make([]coordinates, 0)
	current := coordinates{0, 0}
	result = append(result, current)

	for current.y < len(a.data)-1 {
		current = nav(current)
		result = append(result, current)
	}

	return result
}
