package day08

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func solve1(lines []string) error {
	if instructions, err := parseInstructions(lines); err != nil {
		return err
	} else {
		if res, err := executeJustBeforeInfiniteLoop(instructions); err != nil {
			return err
		} else {
			aoc.PrintSolution(fmt.Sprintf("The acc has the value %d", res))
		}
	}
	return nil
}

func solve2(lines []string) error {
	if instructions, err := parseInstructions(lines); err != nil {
		return err
	} else {
		if res, err := runPermutateInstructions(instructions, func(instructions []instruction) (int, error) {
			return tryExecute(instructions)
		}); err != nil {
			return err
		} else {
			aoc.PrintSolution(fmt.Sprintf("The acc has the value %d", res))
			return nil
		}
	}
}

const nop = "nop"
const acc = "acc"
const jmp = "jmp"

type instruction struct {
	op  string
	val int
}

func parseInstructions(lines []string) ([]instruction, error) {
	result := make([]instruction, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		ins := instruction{op: split[0], val: aoc.ParseInt(split[1])}
		result[i] = ins
	}
	return result, nil
}

type registers struct {
	pc int
	a  int
}

func executeJustBeforeInfiniteLoop(instructions []instruction) (int, error) {
	register := &registers{}
	known := make(map[int]bool)
	if err := execute(register, instructions, func() bool {
		if _, already := known[register.pc]; already {
			return false
		}
		known[register.pc] = true
		return true
	}); err != nil {
		return 0, err
	} else {
		return register.a, nil
	}
}

func tryExecute(instructions []instruction) (int, error) {
	register := &registers{pc: 0, a: 0}
	known := make(map[int]bool)
	infiniteLoop := false
	if err := execute(register, instructions, func() bool {
		if _, already := known[register.pc]; already {
			infiniteLoop = true
			return false
		}
		known[register.pc] = true
		return true
	}); err != nil {
		return 0, err
	} else {
		if infiniteLoop {
			return 0, errors.New(fmt.Sprintf("infinite loop: pc=%d", register.pc))
		}
		return register.a, nil
	}
}

func execute(register *registers,
	instructions []instruction,
	check func() bool) error {

	/*
		fmt.Printf("EXECUTION\n")
		fmt.Printf("=========\n")
		for pc, i := range instructions {
			fmt.Printf("[%d] %s %02d\n", pc, i.op, i.val)
		}
		fmt.Printf("=========\n\n")
	*/

	for check() {
		if !(register.pc < len(instructions)) {
			break
		}
		ins := instructions[register.pc]
		switch ins.op {
		case nop:
			register.pc++
			break
		case jmp:
			register.pc += ins.val
			break
		case acc:
			register.a += ins.val
			register.pc++
			break
		default:
			return errors.New("invalid command")
		}
	}

	return nil
}

func runPermutateInstructions(instructions []instruction, runner func([]instruction) (int, error)) (int, error) {
	flipper := make([]int, 0)
	for i := 0; i < len(instructions); i++ {
		ins := instructions[i]
		if ins.op == jmp || ins.op == nop {
			flipper = append(flipper, i)
		}
	}

	cloner := func(pos int) []instruction {
		clone := make([]instruction, len(instructions))
		for i, ins := range instructions {
			if i == pos {
				var otherOp string
				if ins.op == nop {
					otherOp = jmp
				} else {
					otherOp = nop
				}
				clone[i] = instruction{op: otherOp, val: ins.val}
			} else {
				clone[i] = ins
			}
		}
		return clone
	}

	for idx := 0; idx < len(flipper); idx++ {
		if res, err := runner(instructions); err == nil {
			return res, nil
		}
		if res, err := runner(cloner(idx)); err == nil {
			return res, nil
		}
	}

	return 0, errors.New("no")
}
