package main

import (
	"errors"
	"github.com/knalli/aoc"
	_ "github.com/knalli/aoc2020/day00"
	_ "github.com/knalli/aoc2020/day01"
	_ "github.com/knalli/aoc2020/day02"
	_ "github.com/knalli/aoc2020/day03"
	_ "github.com/knalli/aoc2020/day04"
	_ "github.com/knalli/aoc2020/day05"
	_ "github.com/knalli/aoc2020/day06"
	_ "github.com/knalli/aoc2020/day07"
	_ "github.com/knalli/aoc2020/day08"
	_ "github.com/knalli/aoc2020/day09"
	_ "github.com/knalli/aoc2020/day10"
	_ "github.com/knalli/aoc2020/day11"
	_ "github.com/knalli/aoc2020/day12"
	_ "github.com/knalli/aoc2020/day13"
	_ "github.com/knalli/aoc2020/day14"
	//_ "github.com/knalli/aoc2020/dayXX"
	"os"
	"strconv"
)

func init() {
	aoc.AocYear = 2020
}

func main() {
	if err := invoke(os.Args); err != nil {
		aoc.PrintError(err)
		os.Exit(1)
	}
}

func invoke(args []string) error {
	if len(args) < 2 {
		return errors.New("first argument must be the day (e.g. 1)")
	}
	if day, err := strconv.Atoi(args[1]); err == nil {
		return aoc.Registry.Invoke(day, args[2:])
	} else {
		return err
	}
}
