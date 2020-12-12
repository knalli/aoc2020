package day12

import (
	"github.com/knalli/aoc"
	"github.com/knalli/aoc2020/day11"
)

func ManhattenDistance(p1 *day11.Point, p2 *day11.Point) int {
	return aoc.AbsInt(p2.Y - p1.Y) + aoc.AbsInt(p2.X - p1.X)
}
