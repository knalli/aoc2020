package day14

import (
	"container/list"
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

func solve1(lines []string) error {
	result := run1(lines)
	aoc.PrintSolution(fmt.Sprintf("The sum of all memory values is %d", result))
	return nil
}

func solve2(lines []string) error {
	result := run2(lines)
	aoc.PrintSolution(fmt.Sprintf("The sum of all memory values is %d", result))
	return nil
}

func run1(lines []string) int {
	var mask string
	mem := make(map[int]int)
	for _, line := range lines {
		if line[0:4] == "mask" {
			mask = line[7:]
		} else if line[0:3] == "mem" {
			idx := aoc.ParseInt(line[4:strings.Index(line, "]")])
			val := aoc.ParseInt(line[strings.Index(line, "=")+2:])
			mem[idx] = applyMask(val, mask)
		}
	}

	result := 0
	for _, m := range mem {
		result += m
	}
	return result
}

func run2(lines []string) int {
	var mask string
	mem := make(map[int]int)
	for _, line := range lines {
		if line[0:4] == "mask" {
			mask = line[7:]
		} else if line[0:3] == "mem" {
			idx := aoc.ParseInt(line[4:strings.Index(line, "]")])
			val := aoc.ParseInt(line[strings.Index(line, "=")+2:])
			for _, idx := range maskPermutations(applyMask2(idx, mask)) {
				address := bits2Int(string2Bits(idx))
				mem[address] = val
			}
		}
	}

	result := 0
	for _, m := range mem {
		result += m
	}
	return result
}

func maskPermutations(mask string) []string {
	q := list.New()
	q.PushBack(mask)

	for {
		changed := false

		e := q.Front()
		q.Remove(e)

		current := e.Value.(string)
		if idx := strings.Index(current, "X"); idx > -1 {
			q.PushBack(current[:idx] + "0" + current[idx+1:])
			q.PushBack(current[:idx] + "1" + current[idx+1:])
			changed = true
		} else {
			q.PushBack(current)
		}

		if !changed {
			break
		}
	}

	result := make([]string, 0)
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		result = append(result, e.Value.(string))
	}
	return result
}

func applyMask(val int, mask string) int {
	bits := int2Bits(val, len(mask))
	for i, m := range mask[len(mask)-len(bits):] {
		if m == '0' {
			bits[i] = 0
		} else if m == '1' {
			bits[i] = 1
		}
	}
	return bits2Int(bits)
}

func applyMask2(val int, mask string) string {
	bits := make([]int32, len(mask))
	for i, d := range int2Bits(val, len(mask)) {
		bits[i] = int32(d)
	}
	for i, m := range mask {
		if m == '1' {
			bits[i] = '1'
		} else if m == 'X' {
			bits[i] = 'X'
		} else {
			bits[i] += 48 // conversion ascii
		}
	}
	result := ""
	for _, c := range bits {
		result += string(c)
	}
	return result
}

func int2Bits(val int, len int) []int {
	bits := make([]int, 0)
	r := val
	for i := len - 1; i >= 0; i-- {
		d := int(math.Pow(2, float64(i)))
		if r >= d {
			bits = append(bits, 1)
			r -= d
		} else {
			bits = append(bits, 0)
		}
	}
	return bits
}

func bits2Int(bits []int) int {
	result := 0
	for i, v := range bits {
		n := len(bits) - i - 1
		if v == 1 {
			result += int(math.Pow(2, float64(n)))
		}
	}
	return result
}

func string2Bits(s string) []int {
	result := make([]int, len(s))
	for i, c := range s {
		if c == '1' {
			result[i] = 1
		}
	}
	return result
}
