package day11

import (
	"errors"
	"fmt"
	"io"
)

type IntGrid struct {
	data [][]int

	width  int
	height int
}

type Point struct {
	X int
	Y int
}

func (p Point) Plus(offset *Point) *Point {
	return &Point{X: p.X + offset.X, Y: p.Y + offset.Y}
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

func NewIntGrid(width int, height int) *IntGrid {
	data := make([][]int, height)
	for r := range data {
		data[r] = make([]int, width)
	}
	return &IntGrid{data: data, width: width, height: height}
}

func (g *IntGrid) Clone() *IntGrid {
	clone := NewIntGrid(g.width, g.height)
	for y := range g.data {
		for x, v := range g.data[y] {
			//goland:noinspection GoUnhandledErrorResult
			clone.SetXY(x, y, v)
		}
	}
	return clone
}

func (g *IntGrid) SetXY(x int, y int, v int) error {
	if y < 0 || y >= g.height {
		return errors.New("invalid y")
	}
	if x < 0 || x >= g.width {
		return errors.New("invalid x")
	}
	g.data[y][x] = v
	return nil
}

func (g *IntGrid) GetXY(x int, y int) (int, error) {
	if y < 0 || y >= g.height {
		return 0, errors.New("invalid y")
	}
	if x < 0 || x >= g.width {
		return 0, errors.New("invalid x")
	}
	return g.data[y][x], nil
}

func (g *IntGrid) MustGetXY(x int, y int) int {
	if y < 0 || y >= g.height {
		panic("invalid y")
	}
	if x < 0 || x >= g.width {
		panic("invalid x")
	}
	return g.data[y][x]
}

func (g *IntGrid) Each(r func(p *Point, v int)) {
	for y := range g.data {
		for x, v := range g.data[y] {
			r(NewPoint(x, y), v)
		}
	}
}

func (g *IntGrid) Count(r func(p *Point, v int) bool) int {
	result := 0
	for y := range g.data {
		for x, v := range g.data[y] {
			if r(NewPoint(x, y), v) {
				result++
			}
		}
	}
	return result
}

func (g *IntGrid) CountAdjacents(x int, y int, matcher func(p *Point, v int) bool) int {
	result := 0

	if y > 0 {
		if matcher(NewPoint(x, y-1), g.MustGetXY(x, y-1)) {
			result++
		}
	}
	if x < g.width-1 && y > 0 {
		if matcher(NewPoint(x+1, y-1), g.MustGetXY(x+1, y-1)) {
			result++
		}
	}
	if x < g.width-1 {
		if matcher(NewPoint(x+1, y), g.MustGetXY(x+1, y)) {
			result++
		}
	}
	if x < g.width-1 && y < g.height-1 {
		if matcher(NewPoint(x+1, y+1), g.MustGetXY(x+1, y+1)) {
			result++
		}
	}
	if y < g.height-1 {
		if matcher(NewPoint(x, y+1), g.MustGetXY(x, y+1)) {
			result++
		}
	}
	if x > 0 && y < g.height-1 {
		if matcher(NewPoint(x-1, y+1), g.MustGetXY(x-1, y+1)) {
			result++
		}
	}
	if x > 0 {
		if matcher(NewPoint(x-1, y), g.MustGetXY(x-1, y)) {
			result++
		}
	}
	if x > 0 && y > 0 {
		if matcher(NewPoint(x-1, y-1), g.MustGetXY(x-1, y-1)) {
			result++
		}
	}

	return result
}

func (g *IntGrid) countHelper(p *Point, offset *Point, stopAtFirst bool, filter func(p *Point, v int) bool, matcher func(p *Point, v int) bool) int {

	result := 0
	p = p.Plus(offset)

	for {
		if !(0 <= p.X && p.X < g.width) {
			break
		}
		if !(0 <= p.Y && p.Y < g.height) {
			break
		}

		if matcher(p, g.MustGetXY(p.X, p.Y)) {
			result++
			break
		}
		if !filter(p, g.MustGetXY(p.X, p.Y)) {
			break
		}

		p = p.Plus(offset)
	}

	return result
}

func (g *IntGrid) CountAdjacentVectors(x int, y int, stopAtFirst bool, filter func(p *Point, v int) bool, matcher func(p *Point, v int) bool) int {
	result := 0

	p := NewPoint(x, y)

	if y > 0 {
		offset := NewPoint(0, -1)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if x < g.width-1 && y > 0 {
		offset := NewPoint(1, -1)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if x < g.width-1 {
		offset := NewPoint(1, 0)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if x < g.width-1 && y < g.height-1 {
		offset := NewPoint(1, 1)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if y < g.height-1 {
		offset := NewPoint(0, 1)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if x > 0 && y < g.height-1 {
		offset := NewPoint(-1, 1)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if x > 0 {
		offset := NewPoint(-1, 0)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}
	if x > 0 && y > 0 {
		offset := NewPoint(-1, -1)
		result += g.countHelper(p, offset, stopAtFirst, filter, matcher)
	}

	return result
}

func (g *IntGrid) Print(w io.Writer) {
	for y := range g.data {
		line := ""
		for _, c := range g.data[y] {
			line += string(c)
		}
		fmt.Fprintf(w, "%s\n", line)
	}
}
