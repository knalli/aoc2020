package day05

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	var max *passType
	for _, line := range lines {
		if pass, err := decodeBoardingPass(line); err != nil {
			return err
		} else if max == nil || max.id < pass.id {
			max = pass
		}
	}
	if max == nil {
		return errors.New("no pass found")
	}
	aoc.PrintSolution(fmt.Sprintf("The highest boarding pass ID is %d", max.id))
	return nil
}

func solve2(lines []string) error {
	var passes []*passType
	for _, line := range lines {
		if pass, err := decodeBoardingPass(line); err != nil {
			return err
		} else {
			passes = append(passes, pass)
		}
	}
	countAllEmptySeats := len(findAllEmptySeats(passes))
	if emptySeat, err := findEmptySeat(passes); err != nil {
		return err
	} else {
		aoc.PrintSolution(fmt.Sprintf(
			"The empty seat in row=%d col=%d has boarding pass ID %d (plane has %d empty seats at all)",
			emptySeat.row,
			emptySeat.col,
			passId(emptySeat.row, emptySeat.col),
			countAllEmptySeats,
		))
	}
	return nil
}

type passType struct {
	row int
	col int
	id  int
}

func decodeBoardingPass(str string) (*passType, error) {

	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7

	for i, c := range str {
		if i < 7 {
			if c == 'F' {
				maxRow = maxRow - ((maxRow - minRow + 1) / 2)
			} else if c == 'B' {
				minRow = minRow + (maxRow-minRow)/2 + 1
			} else {
				return nil, errors.New("invalid character: 0-7 must be F or B")
			}
		} else {
			if c == 'L' {
				maxCol = maxCol - ((maxCol - minCol + 1) / 2)
			} else if c == 'R' {
				minCol = minCol + (maxCol-minCol)/2 + 1
			} else {
				return nil, errors.New("invalid character: 8-10 must be L or R")
			}
		}
	}

	if minRow != maxRow || minCol != maxCol {
		return nil, errors.New("invalid row/col found")
	}

	return &passType{row: minRow, col: minCol, id: passId(minRow, minCol)}, nil
}

func passId(row int, col int) int {
	return row*8 + col
}

type seatType struct {
	row int
	col int
}

func findAllEmptySeats(passes []*passType) []seatType {

	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7

	seats := make([][]int, maxRow+1)
	for r := minRow; r <= maxRow; r++ {
		seats[r] = make([]int, maxCol+1)
		for c := minCol; c <= maxCol; c++ {
			seats[r][c] = 0
		}
	}

	for _, pass := range passes {
		seats[pass.row][pass.col] = 1
	}

	result := make([]seatType, 0)
	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if seats[r][c] == 0 {
				result = append(result, seatType{row: r, col: c})
			}
		}
	}

	return result
}

func findEmptySeat(passes []*passType) (*seatType, error) {

	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7

	seats := make([][]int, maxRow+1)
	for r := minRow; r <= maxRow; r++ {
		seats[r] = make([]int, maxCol+1)
		for c := minCol; c <= maxCol; c++ {
			seats[r][c] = 0
		}
	}

	passIds := make(map[int]bool)

	for _, pass := range passes {
		seats[pass.row][pass.col] = 1
		passIds[pass.id] = true
	}

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if seats[r][c] == 0 {
				// id -1
				if _, found := passIds[passId(r, c)-1]; !found {
					continue
				}
				// id +1
				if _, found := passIds[passId(r, c)+1]; !found {
					continue
				}
				return &seatType{row: r, col: c}, nil
			}
		}
	}

	return nil, errors.New("seat not found")
}
