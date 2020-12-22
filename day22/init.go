package day22

import (
	"github.com/knalli/aoc"
)

func init() {
	aoc.Registry.Register(22, main)
}

func main(args []string) error {
	aoc.PrintDayHeader(1, "Day 22")
	if err := step1(args); err != nil {
		return err
	}
	if err := step2(args); err != nil {
		return err
	}
	return nil
}

func step1(args []string) error {
	aoc.PrintStepHeader(1)
	if lines, err := aoc.ReadFileToArray("day22/puzzle1.txt"); err != nil {
		return err
	} else {
		return solve1(lines)
	}
}

func step2(args []string) error {
	aoc.PrintStepHeader(2)
	if lines, err := aoc.ReadFileToArray("day22/puzzle1.txt"); err != nil {
		return err
	} else {
		return solve2(lines)
	}
}

