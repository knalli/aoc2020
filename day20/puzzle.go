package day20

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

const NORTH = 0
const EAST = 1
const SOUTH = 2
const WEST = 3

func solve1(lines []string) error {
	tiles := parseTiles(lines)
	if corners, topLeftTileId, err := findCornerEdges(tiles); err != nil {
		return err
	} else if len(corners) != 4 {
		return errors.New("found not 4 corners")
	} else {
		fmt.Printf("👉 Top left tile id is %d\n", topLeftTileId)
		result := 1
		for _, tileId := range corners {
			result *= tileId
		}
		aoc.PrintSolution(fmt.Sprintf("Product of all corner tile ids is %d\n", result))
	}
	return nil
}

func solve2(lines []string) error {
	tiles := parseTiles(lines)
	_, topLeftTileId, _ := findCornerEdges(tiles)
	assemble(tiles, topLeftTileId, func(land int) {
		aoc.PrintSolution(fmt.Sprintf("%d land left", land))
	})
	return nil
}

func findCornerEdges(tiles []*Tile) ([]int, int, error) {

	corners := make(map[int]bool, 0)
	var bottomLeftTileId int

	for _, tile := range tiles {
		// evaluate corner edges and ensure they are orientated correctly already
		tile.Variants(func(tileVariant *Variant) bool {
			m := make(map[int]int)
			for tileSideId, tileEdge := range tileVariant.Edges {
				for _, other := range tiles {
					if other.Id == tile.Id {
						continue
					}
					if other.Fixed {
						continue
					}
					other.Variants(func(otherVariant *Variant) bool {
						for otherSideId, otherEdge := range otherVariant.Edges {
							if otherSideId == (2+(tileSideId))%4 && tileEdge == otherEdge {
								m[tileSideId] = m[tileSideId] + 1
							}
						}
						return false
					})
				}
			}
			if len(m) == 2 {
				corners[tile.Id] = true
				if bottomLeftTileId == 0 && m[SOUTH] > 0 && m[EAST] > 0 {
					bottomLeftTileId = tile.Id
					return true
				}
			}
			return false
		})
	}
	cornersList := make([]int, 0)
	for id := range corners {
		cornersList = append(cornersList, id)
	}
	return cornersList, bottomLeftTileId, nil
}

func assemble(tiles []*Tile, startTileId int, answer func(land int)) {
	tileById := func(id int) *Tile {
		for _, t := range tiles {
			if t.Id == id {
				return t
			}
		}
		panic("invalid tileId")
	}

	startTile := tileById(startTileId)
	positions := make(map[string]int)
	width := int(math.Sqrt(float64(len(tiles))))

	{
		setPosition := func(x int, y int, tileId int) {
			//fmt.Printf("Pos %d/%d => %d\n", x, y, tileId)
			positions[fmt.Sprintf("%d/%d", x, y)] = tileId
		}

		findNeighbour := func(tile *Tile, tileSide int) *Tile {
			otherSide := (2 + tileSide) % 4
			for _, other := range tiles {
				if tile.Id == other.Id {
					continue
				}
				skip := false
				for _, tileIdWithPosition := range positions {
					if tileIdWithPosition == other.Id {
						skip = true
						break
					}
				}
				if skip {
					continue
				}
				found := other.Variants(func(otherVariant *Variant) bool {
					return tile.Grid.getEdges()[tileSide] == otherVariant.Edges[otherSide]
				})
				if found != nil {
					other.Fixed = true
					return other
				}
			}
			panic("found no neighbour")
		}

		startTile.Grid.PrintToString()
		mostRecentTile := startTile
		mostRecentRowHeader := startTile
		for col := 0; col < width; col++ {
			for row := 0; row < width; row++ {
				if col == 0 && row == 0 {
					// top left
					setPosition(row, col, mostRecentTile.Id)
				} else if row == 0 {
					// find via next south
					mostRecentRowHeader = findNeighbour(mostRecentRowHeader, SOUTH)
					mostRecentTile = mostRecentRowHeader
					setPosition(row, col, mostRecentRowHeader.Id)
				} else {
					mostRecentTile = findNeighbour(mostRecentTile, EAST)
					setPosition(row, col, mostRecentTile.Id)
				}
			}
		}
	}

	// build super tile (aka the "image")
	tileSize := startTile.Grid.width - 2
	superTile := Tile{
		Id:   0,
		Grid: NewGrid(width*tileSize, width*tileSize),
	}

	for y := 0; y < superTile.Grid.width; y++ {
		for x := 0; x < superTile.Grid.width; x++ {
			superTile.Grid.SetXY(x, y, 'X')
		}
	}

	for row := 0; row < width; row++ {
		for col := 0; col < width; col++ {
			tile := tileById(positions[fmt.Sprintf("%d/%d", col, row)])
			for y := 1; y < tile.Grid.height-1; y++ {
				for x := 1; x < tile.Grid.width-1; x++ {
					sx := (col * tileSize) + (x - 1)
					sy := (row * tileSize) + (y - 1)
					superTile.Grid.SetXY(
						sx,
						sy,
						tile.Grid.data[y][x],
					)
				}
			}
		}
	}
	//superTile.Grid.Print(os.Stdout)

	mask := make([][]int, 0)
	for _, search := range []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	} {
		r := make([]int, 0)
		for i, c := range search {
			if c == '#' {
				r = append(r, i)
			}
		}
		mask = append(mask, r)
	}

	for _, flipped := range []bool{true, false} {
		for _, rotation := range []int{0, 1, 2, 3} {
			clone := superTile.Grid.Clone()
			for r := 0; r < rotation; r++ {
				clone.RotateRight()
			}
			if flipped {
				clone.FlipY()
			}
			found, land := findMonster(clone, mask)
			if found {
				answer(land)
			}
		}
	}
}

func findMonster(g *CharGrid, mask [][]int) (bool, int) {

	height := len(mask)
	var width int
	for _, maskRow := range mask {
		for _, s := range maskRow {
			width = aoc.MaxInt(width, s)
		}
	}

	count := 0
	for y := 0; y < g.height-height; y++ {
		for x := 0; x < g.width-width; x++ {
			ok := true
			for dy := range mask {
				for _, dx := range mask[dy] {
					if g.data[y+dy][x+dx] != '#' {
						ok = false
						break
					}
				}
				if !ok {
					break
				}
			}
			if ok {
				count++
				for dy := range mask {
					for _, dx := range mask[dy] {
						g.data[y+dy][x+dx] = '0'
					}
				}
			}
		}
	}

	land := 0
	if count > 0 {
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				if g.data[y][x] == '#' {
					land++
				}
			}
		}
	}

	return count > 0, land
}

func parseTiles(lines []string) []*Tile {
	tiles := make([]*Tile, 0)

	var grid *CharGrid
	row := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		} else if strings.Index(line, "Tile ") == 0 {
			grid = NewGrid(10, 10)
			tiles = append(tiles, &Tile{Id: aoc.ParseInt(line[5 : len(line)-1]), Grid: grid, UsedEdges: make(map[int]bool)})
			row = 0
		} else {
			for x, c := range line {
				if c == '#' || c == '.' {
					grid.SetXY(x, row, c)
				} else {
					panic("invalid value")
				}
			}
			row++
		}
	}

	return tiles
}

type Context struct {
	Tiles       map[int]Params
	Connections map[string]bool
	Positions   map[string]int
}

func NewContext() Context {
	return Context{
		Tiles:       make(map[int]Params),
		Connections: make(map[string]bool),
		Positions:   make(map[string]int),
	}
}

func (c *Context) Add(tileId int, rotation int, flippedX bool, flippedY bool, tileEdges []string, positionX int, positionY int) {
	c.Tiles[tileId] = Params{
		PositionX: positionX,
		PositionY: positionY,
		Rotation:  rotation,
		FlippedX:  flippedX,
		FlippedY:  flippedY,
		Edges:     tileEdges,
	}
	c.Positions[fmt.Sprintf("%d/%d", positionX, positionY)] = tileId
}

func (c *Context) Clone() *Context {
	clone := Context{}
	clone.Connections = make(map[string]bool)
	for s, v := range c.Connections {
		clone.Connections[s] = v
	}
	clone.Tiles = make(map[int]Params)
	for i, v := range c.Tiles {
		clone.Tiles[i] = Params{
			PositionX: v.PositionX,
			PositionY: v.PositionY,
			Rotation:  v.Rotation,
			FlippedX:  v.FlippedX,
			FlippedY:  v.FlippedY,
			Edges:     v.Edges,
		}
	}
	clone.Positions = make(map[string]int)
	for p, v := range c.Positions {
		clone.Positions[p] = v
	}
	return &clone
}

type Params struct {
	PositionX int
	PositionY int
	Rotation  int
	FlippedX  bool
	FlippedY  bool
	Edges     []string
}

func rotateRightGridEdges(input []string) []string {
	return []string{
		reverseString(input[WEST]),
		input[NORTH],
		reverseString(input[EAST]),
		input[SOUTH],
	}
}

func flipGridEdges(input []string) []string {
	return []string{
		reverseString(input[NORTH]),
		input[WEST],
		reverseString(input[SOUTH]),
		input[EAST],
	}
}

func reverseString(s string) string {
	r := make([]uint8, 0)
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return string(r)
}
