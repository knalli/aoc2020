package day00

import (
	"fmt"
	"github.com/knalli/aoc"
)

func init() {
	aoc.Registry.Register(0, main)
}

func main(args []string) error {
	aoc.PrintDayHeader(0, "It works")

	for i := range args {
		fmt.Printf("args[%d] = %s\n", i, args[i])
	}

	return nil
}

func greet(name string) string {
	return "Hello " + name
}
