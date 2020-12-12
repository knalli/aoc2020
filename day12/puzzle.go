package day12

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"github.com/knalli/aoc2020/day11"
)

func solve1(lines []string) error {
	if instructions, err := parseInstructions(lines); err != nil {
		return err
	} else {
		startPos := day11.NewPoint(0, 0)
		if endPos, err := followInstructions1(startPos, EAST, instructions); err != nil {
			return err
		} else {
			aoc.PrintSolution(fmt.Sprintf("The end position is %d/%d with d = %d", endPos.X, endPos.Y, ManhattenDistance(startPos, endPos)))
		}
	}
	return nil
}

func solve2(lines []string) error {
	if instructions, err := parseInstructions(lines); err != nil {
		return err
	} else {
		startPos := day11.NewPoint(0, 0)
		waypoint := day11.NewPoint(10, 1)
		if endPos, err := followInstructions2(startPos, waypoint, instructions); err != nil {
			return err
		} else {
			aoc.PrintSolution(fmt.Sprintf("The end position is %d/%d with d = %d", endPos.X, endPos.Y, ManhattenDistance(startPos, endPos)))
		}
	}
	return nil
}

type Direction int32

const NORTH = Direction('N')
const EAST = Direction('E')
const SOUTH = Direction('S')
const WEST = Direction('W')
const LEFT = Direction('L')
const RIGHT = Direction('R')
const FORWARD = Direction('F')

type Instruction struct {
	Direction Direction
	Amount    int
}

func parseInstructions(lines []string) ([]Instruction, error) {
	result := make([]Instruction, len(lines))
	for i, line := range lines {
		result[i] = Instruction{Direction: Direction(line[0]), Amount: aoc.ParseInt(line[1:])}
	}
	return result, nil
}
func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func followInstructions1(pos *day11.Point, dir Direction, instructions []Instruction) (*day11.Point, error) {
	dirOrder := []Direction{NORTH, EAST, SOUTH, WEST}

	for _, ins := range instructions {
		switch ins.Direction {
		case NORTH:
			pos = pos.Plus(&day11.Point{X: 0, Y: ins.Amount})
			break
		case EAST:
			pos = pos.Plus(&day11.Point{X: ins.Amount, Y: 0})
			break
		case SOUTH:
			pos = pos.Plus(&day11.Point{X: 0, Y: -ins.Amount})
			break
		case WEST:
			pos = pos.Plus(&day11.Point{X: -ins.Amount, Y: 0})
			break
		case LEFT:
			currentDirectionIndex := sliceIndex(len(dirOrder), func(i int) bool {
				return dirOrder[i] == dir
			})
			dir = dirOrder[(currentDirectionIndex-(ins.Amount/90)+4)%4]
			break
		case RIGHT:
			currentDirectionIndex := sliceIndex(len(dirOrder), func(i int) bool {
				return dirOrder[i] == dir
			})
			dir = dirOrder[(currentDirectionIndex+(ins.Amount/90)+4)%4]
			break
		case FORWARD:
			switch dir {
			case NORTH:
				pos = pos.Plus(&day11.Point{X: 0, Y: ins.Amount})
				break
			case EAST:
				pos = pos.Plus(&day11.Point{X: ins.Amount, Y: 0})
				break
			case SOUTH:
				pos = pos.Plus(&day11.Point{X: 0, Y: -ins.Amount})
				break
			case WEST:
				pos = pos.Plus(&day11.Point{X: -ins.Amount, Y: 0})
				break
			default:
				return nil, errors.New("invalid instruction")
			}
			break
		default:
			return nil, errors.New("invalid instruction")
		}

		//fmt.Printf("Position is: %d/%d [%c]\n", pos.X, pos.Y, dir)
	}

	return pos, nil
}

func followInstructions2(pos *day11.Point, wpOffset *day11.Point, instructions []Instruction) (*day11.Point, error) {

	for _, ins := range instructions {
		switch ins.Direction {
		case NORTH:
			wpOffset = wpOffset.Plus(&day11.Point{X: 0, Y: ins.Amount})
			break
		case EAST:
			wpOffset = wpOffset.Plus(&day11.Point{X: ins.Amount, Y: 0})
			break
		case SOUTH:
			wpOffset = wpOffset.Plus(&day11.Point{X: 0, Y: -ins.Amount})
			break
		case WEST:
			wpOffset = wpOffset.Plus(&day11.Point{X: -ins.Amount, Y: 0})
			break
		case LEFT:
			for i := 0; i < ins.Amount/90; i++ {
				wpOffset = day11.NewPoint(-wpOffset.Y, wpOffset.X)
			}
			break
		case RIGHT:
			for i := 0; i < ins.Amount/90; i++ {
				wpOffset = day11.NewPoint(wpOffset.Y, -wpOffset.X)
			}
			break
		case FORWARD:
			for i := 0; i < ins.Amount; i++ {
				pos = pos.Plus(wpOffset)
			}
		default:
			return nil, errors.New("invalid instruction")
		}

		//fmt.Printf("%c%d: Position: %d/%d [%c], Waypoint offset: %d/%d\n", ins.Direction, ins.Amount, pos.X, pos.Y, dir, wpOffset.X, wpOffset.Y)
	}

	return pos, nil
}
