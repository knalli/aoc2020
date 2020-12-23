package day23

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

func solve1(lines []string) error {
	moves := 100
	result, _ := play(
		moves,
		&Context{cups: parse(lines[0]), picks: make([]int, 3)},
		true,
		false,
		false,
	)
	aoc.PrintSolution(fmt.Sprintf("After %d moves, cups order is '%s'", moves, result))
	return nil
}

func solve2(lines []string) error {

	moves := 10000000
	cups := parse(lines[0])
	maxValue := math.MinInt64
	cups.Each(func(e *aoc.LinkElement) {
		v := e.Value().(int)
		maxValue = aoc.MaxInt(maxValue, v)
	})
	for i := maxValue + 1; i <= 1000000; i++ {
		cups.Add(i)
	}
	_, result := play(
		moves,
		&Context{cups: cups, picks: make([]int, 3)},
		false,
		true,
		false,
	)
	aoc.PrintSolution(fmt.Sprintf("After %d moves, r0=%d, r1=%d and r0*r1=%d", moves, result[0], result[1], result[0]*result[1]))
	return nil
}

type Context struct {
	move        int
	cups        *aoc.LinkedList
	current     *aoc.LinkElement
	picks       []int
	destination *aoc.LinkElement
}

func parse(line string) *aoc.LinkedList {
	l := aoc.NewLinkedList(true)
	for _, c := range line {
		l.Add(int(c - 48))
	}
	return l
}

func play(rounds int, ctx *Context, part1 bool, part2 bool, debug bool) (string, []int) {

	ctx.move = 1
	ctx.current = ctx.cups.Front()

	minValue := math.MaxInt64
	maxValue := math.MinInt64
	ctx.cups.Each(func(e *aoc.LinkElement) {
		v := e.Value().(int)
		minValue = aoc.MinInt(minValue, v)
		maxValue = aoc.MaxInt(maxValue, v)
	})

	// Relevant for part2 (improve performance drastically)
	links := make(map[int]*aoc.LinkElement)
	ctx.cups.Each(func(e *aoc.LinkElement) {
		links[e.Value().(int)] = e
	})

	for ; ctx.move <= rounds; ctx.move++ {
		debugMessages := make([]string, 0)
		if debug {
			fmt.Printf("\n-- move %d --\n", ctx.move)
			fmt.Printf("cups: %s\n", ctx.cups.ToString(func(e *aoc.LinkElement) string {
				if e == ctx.current {
					return fmt.Sprintf("(%d)", e.Value().(int))
				} else {
					return fmt.Sprintf(" %d ", e.Value().(int))
				}
			}))
		}
		// collect next 3
		{
			for i := 0; i < 3; i++ {
				next := ctx.current.Next()
				ctx.picks[i] = next.Value().(int)
				ctx.cups.Remove(next)
			}
			if debug {
				picksS := make([]string, len(ctx.picks))
				for i, v := range ctx.picks {
					picksS[i] = fmt.Sprintf("%d", v)
				}
				fmt.Printf("pick up: %s\n", strings.Join(picksS, ", "))
			}
		}
		// find dest
		for search := ctx.current.Value().(int) - 1; ; {
			if search < minValue {
				search = maxValue
				continue
			}
			taken := false
			for i := 0; i < len(ctx.picks); i++ {
				if ctx.picks[i] == search {
					taken = true
					break
				}
			}
			if taken {
				search--
				continue
			}
			ctx.destination = links[search]
			if ctx.destination == nil {
				panic("destination not found")
			}
			break
		}
		if debug {
			debugMessages = append(
				debugMessages,
				fmt.Sprintf("destination: %d", ctx.destination.Value().(int)),
			)
			fmt.Println(strings.Join(debugMessages, "\n"))
		}
		{
			e := ctx.destination
			for i := 0; i < 3; i++ {
				e = ctx.cups.AddAfter(e, ctx.picks[i])
				links[ctx.picks[i]] = e // don't forget updating new ref
			}
		}

		ctx.current = ctx.current.Next()
	}

	if debug {
		fmt.Println("\n-- final --")
		fmt.Printf("cups: %s\n", ctx.cups.ToString(func(e *aoc.LinkElement) string {
			if e == ctx.current {
				return fmt.Sprintf("(%d)", e.Value().(int))
			} else {
				return fmt.Sprintf(" %d ", e.Value().(int))
			}
		}))
	}

	labelsAsString := ""
	firstLabelPair := make([]int, 0)
	if part1 {
		ctx.cups.EachStartAt(links[1], func(e *aoc.LinkElement) {
			labelsAsString += fmt.Sprintf("%d", e.Value().(int))
		})
		labelsAsString = labelsAsString[1:]
	}
	if part2 {
		firstLabelPair = []int{
			links[1].Next().Value().(int),
			links[1].Next().Next().Value().(int),
		}
	}
	return labelsAsString, firstLabelPair
}
