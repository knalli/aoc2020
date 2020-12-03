package day03

import "errors"

const EMTPY = "."
const FILLED = "#"

const D_EMPTY = "⬜️"
const D_TREE = "🌲"

type coordinates struct {
	x int
	y int
}

type area struct {
	data     [][]string
	minWidth int
}

func (a *area) toString(pretty bool) string {
	var renderEmpty string
	var renderFilled string
	if pretty {
		renderEmpty = D_EMPTY
		renderFilled = D_TREE
	} else {
		renderEmpty = EMTPY
		renderFilled = FILLED
	}

	out := ""
	for _, line := range a.data {
		out = a.renderInternalLine(line, out, renderFilled, renderEmpty, 0)
		out += "\n"
	}

	return out
}

func (a *area) renderInternalLine(line []string, result string, filledString string, emptyString string, offset int) string {
	for _, s := range line {
		if s == FILLED {
			result += filledString
		} else {
			result += emptyString
		}
	}
	lineWidth := len(line)
	if a.minWidth-offset > lineWidth {
		return a.renderInternalLine(line, result, filledString, emptyString, offset+lineWidth)
	}
	return result
}

func (a *area) get(c coordinates) (string, error) {
	if c.y >= len(a.data) {
		return "", errors.New("invalid coordinate y")
	}

	lineWidth := len(a.data[c.y])
	for c.x > a.minWidth {
		a.minWidth += lineWidth
	}
	return a.data[c.y][c.x%lineWidth], nil
}

func parseArea(lines []string) *area {
	data := make([][]string, len(lines))
	for i, line := range lines {
		data[i] = make([]string, len(line))
		for j, chr := range line {
			s := string(chr)
			data[i][j] = s
		}
	}
	return &area{data: data, minWidth: len(data[0])}
}
