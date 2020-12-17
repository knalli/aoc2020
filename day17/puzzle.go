package day17

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	pocket := simulate1(lines, 6)
	aoc.PrintSolution(fmt.Sprintf("After 6 cycles, there are %d cubes active", countActive2(pocket)))
	return nil
}

func solve2(lines []string) error {
	pocket := simulate2(lines, 6)
	aoc.PrintSolution(fmt.Sprintf("After 6 cycles, there are %d cubes active", countActive2(pocket)))
	return nil
}

type Cube struct {
	value      int
	neighbours int
}

func simulate1(lines []string, times int) map[string]Cube {
	pocket := make(map[string]Cube)
	for y, line := range lines {
		for x, c := range line {
			var v int
			if c == '#' {
				v = 1
			}
			pocket[fmt.Sprintf("%d,%d,%d", x, y, 0)] = Cube{value: v}
		}
	}

	for i := 0; i < times; i++ {
		updateAdjacentCounters1(pocket)
		runCycle(pocket)
	}

	return pocket
}

func simulate2(lines []string, times int) map[string]Cube {
	pocket := make(map[string]Cube)
	for y, line := range lines {
		for x, c := range line {
			var v int
			if c == '#' {
				v = 1
			}
			pocket[fmt.Sprintf("%d,%d,%d,%d", x, y, 0, 9)] = Cube{value: v}
		}
	}

	for i := 0; i < times; i++ {
		updateAdjacentCounters2(pocket)
		runCycle(pocket)
	}

	return pocket
}

func updateAdjacentCounters1(pocket map[string]Cube) {
	for id, cube := range pocket {

		if cube.value == 0 {
			continue // just created
		}

		var x, y, z int
		//goland:noinspection GoUnhandledErrorResult
		fmt.Sscanf(id, "%d,%d,%d", &x, &y, &z)

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				for dz := -1; dz <= 1; dz++ {
					adjacentId := fmt.Sprintf("%d,%d,%d", x+dx, y+dy, z+dz)
					if adjacentId != id {
						adjacentCube := pocket[adjacentId] // beware go returns an empty struct if not found
						adjacentCube.neighbours++
						pocket[adjacentId] = adjacentCube
					}
				}
			}
		}
	}
}

func updateAdjacentCounters2(pocket map[string]Cube) {
	for id, cube := range pocket {

		if cube.value == 0 {
			continue // just created
		}

		var x, y, z, w int
		//goland:noinspection GoUnhandledErrorResult
		fmt.Sscanf(id, "%d,%d,%d,%d", &x, &y, &z, &w)

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				for dz := -1; dz <= 1; dz++ {
					for dw := -1; dw <= 1; dw++ {
						adjacentId := fmt.Sprintf("%d,%d,%d,%d", x+dx, y+dy, z+dz, w+dw)
						if adjacentId != id {
							adjacentCube := pocket[adjacentId] // beware go returns an empty struct if not found
							adjacentCube.neighbours++
							pocket[adjacentId] = adjacentCube
						}
					}
				}
			}
		}
	}
}

func runCycle(pocket map[string]Cube) {
	for id, cube := range pocket {
		if cube.value == 1 && !(cube.neighbours == 2 || cube.neighbours == 3) {
			delete(pocket, id)
		} else if cube.value == 0 && cube.neighbours != 3 {
			delete(pocket, id)
		} else {
			pocket[id] = Cube{1, 0}
		}
	}
}

func countActive2(pocket map[string]Cube) int {
	result := 0
	for _, cube := range pocket {
		if cube.value == 1 {
			result++
		}
	}
	return result
}
