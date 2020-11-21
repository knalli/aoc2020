package main

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"github.com/knalli/aoc2020/day00"
	"os"
	"strconv"
)

func main() {
	registerAll()
	if err := invoke(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerAll() {
	aoc.AocYear = 2020
	aoc.Registry.Register(00, day00.Call)
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
