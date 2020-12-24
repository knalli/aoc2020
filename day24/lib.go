package day24

import "github.com/knalli/aoc2020/day11"

type Cube struct {
	X int
	Y int
	Z int
}

func (c *Cube) Add(other Cube) Cube {
	return Cube{
		X: c.X + other.X,
		Y: c.Y + other.Y,
		Z: c.Z + other.Z,
	}
}

func (c *Cube) Neighbor(direction HexDirection) Cube {
	return c.Add(directions[direction])
}

func (c *Cube) Neighbors() []Cube {
	result := make([]Cube, len(directions))
	for i := 0; i < len(directions); i++ {
		result[i] = c.Add(directions[i])
	}
	return result
}

func (c *Cube) EachNeighbor(iter func(n Cube)) {
	for i := 0; i < len(directions); i++ {
		iter(c.Add(directions[i]))
	}
}

func (c *Cube) Point_oddr() day11.Point {
	col := c.X + (c.Z-(c.Z&1))/2
	row := c.Z
	return day11.Point{X: col, Y: row}
}

func (c *Cube) Point_evenr() day11.Point {
	col := c.X + (c.Z+(c.Z&1))/2
	row := c.Z
	return day11.Point{X: col, Y: row}
}

type HexDirection int

const EAST HexDirection = 0
const SOUTHEAST HexDirection = 1
const SOUTHWEST HexDirection = 2
const WEST HexDirection = 3
const NORTHWEST HexDirection = 4
const NORTHEAST HexDirection = 5

var directions = []Cube{
	Cube{1, -1, 0},
	Cube{1, 0, -1},
	Cube{0, 1, -1},
	Cube{-1, 1, 0},
	Cube{-1, 0, 1},
	Cube{0, -1, 1},
}
