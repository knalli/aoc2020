package day00

import (
	"fmt"
	"github.com/knalli/aoc"
)

func Call(args []string) error {
	aoc.PrintDayHeader(0, "It works")

	for i := range args {
		fmt.Printf("args[%d] = %s\n", i, args[i])
	}

	return nil
}
