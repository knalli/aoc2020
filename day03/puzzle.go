package day03

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	area := parseArea(lines)
	path := runSlope(area, func(c coordinates) coordinates {
		return coordinates{x: c.x + 3, y: c.y + 1}
	})
	if count, err := countFilled(path, area); err != nil {
		return err
	} else {
		//fmt.Println(area.toString(false))
		aoc.PrintSolution(fmt.Sprintf("Tobogann will encounter %d trees", count))
	}
	return nil
}

func solve2(lines []string) error {
	area := parseArea(lines)

	paths := runSlopes(
		area,
		func(c coordinates) coordinates {
			return coordinates{x: c.x + 1, y: c.y + 1}
		},
		func(c coordinates) coordinates {
			return coordinates{x: c.x + 3, y: c.y + 1}
		},
		func(c coordinates) coordinates {
			return coordinates{x: c.x + 5, y: c.y + 1}
		},
		func(c coordinates) coordinates {
			return coordinates{x: c.x + 7, y: c.y + 1}
		},
		func(c coordinates) coordinates {
			return coordinates{x: c.x + 1, y: c.y + 2}
		},
	)
	result := 1
	for _, way := range paths {
		if count, err := countFilled(way, area); err != nil {
			return err
		} else {
			result *= count
		}
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

func runSlopes(a *area, navs ...navigator) [][]coordinates {
	result := make([][]coordinates, len(navs))
	for i, nav := range navs {
		result[i] = runSlope(a, nav)
	}
	return result
}
