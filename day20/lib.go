package day20

type Tile struct {
	Id        int
	Grid      *CharGrid
	Rotation  int
	Flipped   bool
	UsedEdges map[int]bool
	Fixed     bool
	PositionX int
	PositionY int
}

type Variant struct {
	Rotation int
	Flipped  bool
	Edges    []string
}

func (t *Tile) Variants(f func(variant *Variant) bool) *Variant {
	edges := t.Grid.getEdges()

	test := func(rotation int, flipped bool) *Variant {
		e := edges
		for r := 0; r < rotation; r++ {
			e = rotateRightGridEdges(e)
		}
		if flipped {
			e = flipGridEdges(e)
		}
		variant := &Variant{
			Rotation: rotation,
			Flipped:  flipped,
			Edges:    e,
		}
		if f(variant) {
			//fmt.Printf("Tile %d will be transformed: rotation=%d, flipped=%v\n", t.Id, rotation, flipped)
			for r := 0; r < rotation; r++ {
				t.Grid.RotateRight()
			}
			if flipped {
				t.Grid.FlipY()
			}
			return variant
		}
		return nil
	}

	for _, flipped := range []bool{true, false} {
		for _, rotation := range []int{0, 1, 2, 3} {
			if result := test(rotation, flipped); result != nil {
				return result
			}
		}
	}
	return nil
}
