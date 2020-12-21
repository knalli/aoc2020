package day20

import (
	"bytes"
	"fmt"
	"io"
)

type CharGrid struct {
	data   [][]int32
	height int
	width  int
}

func NewGrid(height int, width int) *CharGrid {
	data := buildEmptyData(height, width)
	return &CharGrid{data: data, height: height, width: width}
}

func (g *CharGrid) Clone() *CharGrid {
	clone := NewGrid(g.height, g.width)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			clone.SetXY(x, y, g.data[y][x])
		}
	}
	return clone
}

func buildEmptyData(height int, width int) [][]int32 {
	data := make([][]int32, height)
	for i := range data {
		data[i] = make([]int32, width)
	}
	return data
}

func (g *CharGrid) RotateRight() {
	height := g.width
	width := g.height
	next := buildEmptyData(height, width)

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			next[x][y] = g.data[y][x]
		}
	}
	for y := 0; y < len(next); y++ {
		n := make([]int32, g.width)
		for i, c := range reverseString(string(next[y])) {
			n[i] = c
		}
		next[y] = n
	}

	g.data = next
	g.height = height
	g.width = width
}

func (g *CharGrid) FlipY() {
	next := buildEmptyData(g.height, g.width)

	for y := 0; y < g.height; y++ {
		n := make([]int32, g.width)
		for i, c := range reverseString(string(g.data[y])) {
			n[i] = c
		}
		next[y] = n
	}

	g.data = next
}

func (g *CharGrid) getEdges() []string {
	return []string{
		g.getNorthEdge(),
		g.getEastEdge(),
		g.getSouthEdge(),
		g.getWestEdge(),
	}
}

func (g *CharGrid) getNorthEdge() string {
	return string(g.data[0])
}

func (g *CharGrid) getSouthEdge() string {
	return string(g.data[g.height-1])
}

func (g *CharGrid) getWestEdge() string {
	s := ""
	for y := range g.data {
		s += string(g.data[y][0])
	}
	return s
}

func (g *CharGrid) getEastEdge() string {
	s := ""
	for y := range g.data {
		s += string(g.data[y][g.width-1])
	}
	return s
}

func (g *CharGrid) SetXY(x int, y int, v int32) {
	g.data[y][x] = v
}

func (g *CharGrid) Print(w io.Writer) {
	s := ""
	for y := range g.data {
		for _, c := range g.data[y] {
			s += string(c)
		}
		s += "\n"
	}
	fmt.Fprintln(w, s)
}

func (g *CharGrid) PrintToString() string {
	buf := new(bytes.Buffer)
	g.Print(buf)
	return buf.String()
}
