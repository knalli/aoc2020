package day24

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	tiles := initialize(lines)
	aoc.PrintSolution(fmt.Sprintf("flipped to black = %d", len(tiles)))
	return nil
}

func solve2(lines []string) error {
	tiles := dailyFlips(initialize(lines), 100)
	aoc.PrintSolution(fmt.Sprintf("flipped to black = %d", len(tiles)))
	return nil
}

func initialize(lines []string) map[Cube]bool {
	tiles := make(map[Cube]bool)
	start := Cube{}
	for _, line := range lines {
		c := findDestination(start, line)
		if tiles[c] {
			delete(tiles, c)
		} else {
			tiles[c] = true
		}
	}
	return tiles
}

func dailyFlips(tiles map[Cube]bool, days int) map[Cube]bool {
	for day := 0; day < days; day++ {

		// collect all relevant (incl. neighbors because 2nd rule "any white")
		scope := make([]Cube, 0)
		for c := range tiles {
			scope = append(scope, c.Neighbors()...)
		}

		nextTiles := make(map[Cube]bool)

		for _, c := range distinctCubes(scope) {
			flipped := tiles[c]
			neighborsFlipped := 0
			c.EachNeighbor(func(n Cube) {
				if tiles[n] {
					neighborsFlipped++
				}
			})
			if flipped && (neighborsFlipped == 1 || neighborsFlipped == 2) {
				nextTiles[c] = true // stay with flipped
			}
			if !flipped && neighborsFlipped == 2 {
				nextTiles[c] = true // flip
			}
			// otherwise: no change of flip to white
		}
		tiles = nextTiles
		//fmt.Printf("Day %d: %d\n", day+1, len(tiles))
	}
	return tiles
}

func findDestination(start Cube, line string) Cube {
	dest := start
	for i := 0; i < len(line); i++ {
		if line[i] == 'e' {
			dest = dest.Neighbor(EAST)
		} else if line[i] == 'w' {
			dest = dest.Neighbor(WEST)
		} else {
			switch line[i : i+2] {
			case "se":
				dest = dest.Neighbor(SOUTHEAST)
				i++
				break
			case "sw":
				dest = dest.Neighbor(SOUTHWEST)
				i++
				break
			case "ne":
				dest = dest.Neighbor(NORTHEAST)
				i++
				break
			case "nw":
				dest = dest.Neighbor(NORTHWEST)
				i++
				break
			default:
				panic("invalid direction")
			}
		}
	}
	return dest
}

func distinctCubes(slice []Cube) []Cube {
	keys := make(map[Cube]bool)
	var list []Cube
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
