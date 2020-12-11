package day11

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	if grid, err := parseGrid(lines); err != nil {
		return err
	} else {
		round := 0
		changes := 0
		for {
			changes, grid = doPlacements1(grid)
			round++
			if changes == 0 {
				break
			}
		}
		occupiedSeats := grid.Count(func(p *Point, v int) bool {
			return v == SEAT
		})
		aoc.PrintSolution(fmt.Sprintf("%d occupied seats", occupiedSeats))
	}
	return nil
}

func solve2(lines []string) error {
	if grid, err := parseGrid(lines); err != nil {
		return err
	} else {
		round := 0
		changes := 0
		//grid.Print(os.Stdout)
		for {
			changes, grid = doPlacements2(grid)
			//fmt.Println()
			//grid.Print(os.Stdout)
			round++
			if changes == 0 {
				break
			}
		}
		occupiedSeats := grid.Count(func(p *Point, v int) bool {
			return v == SEAT
		})
		aoc.PrintSolution(fmt.Sprintf("%d occupied seats", occupiedSeats))
	}
	return nil
}

const EMPTY = 'L'
const FLOOR = '.'
const SEAT = '#'

func parseGrid(lines []string) (*IntGrid, error) {

	height := len(lines)
	if height == 0 {
		return nil, errors.New("empty rows")
	}

	width := len(lines[0])
	if height == 0 {
		return nil, errors.New("empty columns")
	}

	grid := NewIntGrid(width, height)

	for y, line := range lines {
		for x, c := range line {
			if c == EMPTY || c == FLOOR || c == SEAT {
				//goland:noinspection GoUnhandledErrorResult
				grid.SetXY(x, y, int(c))
			} else {
				return nil, errors.New("invalid character found")
			}
		}
	}

	return grid, nil
}

func doPlacements1(grid *IntGrid) (int, *IntGrid) {
	next := grid.Clone()
	changes := 0
	grid.Each(func(p *Point, v int) {
		occupiedAdjacents := grid.CountAdjacents(p.X, p.Y, func(p *Point, v int) bool {
			return v == SEAT
		})
		if v == EMPTY && occupiedAdjacents == 0 {
			//goland:noinspection GoUnhandledErrorResult
			next.SetXY(p.X, p.Y, SEAT)
			changes++
		} else if v == SEAT && occupiedAdjacents >= 4 {
			//goland:noinspection GoUnhandledErrorResult
			next.SetXY(p.X, p.Y, EMPTY)
			changes++
		}
	})

	return changes, next
}

func doPlacements2(grid *IntGrid) (int, *IntGrid) {
	next := grid.Clone()
	changes := 0
	grid.Each(func(p *Point, v int) {
		occupiedAdjacents := grid.CountAdjacentVectors(p.X, p.Y, true, func(p *Point, v int) bool {
			return v == FLOOR
		}, func(p *Point, v int) bool {
			return v == SEAT
		})
		if v == EMPTY && occupiedAdjacents == 0 {
			//goland:noinspection GoUnhandledErrorResult
			next.SetXY(p.X, p.Y, SEAT)
			changes++
		} else if v == SEAT && occupiedAdjacents >= 5 {
			//goland:noinspection GoUnhandledErrorResult
			next.SetXY(p.X, p.Y, EMPTY)
			changes++
		}
	})

	return changes, next
}
