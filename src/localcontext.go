package src

/// Pixel position type
type PixelPos struct {
	Row,
	Col int
}

/// Pixel mask type
type ContextMask []PixelPos

/// Conext table type
type ContextTable map[PixelPos]ContextMask

/// Accessor
func (c *ContextTable) GetMask(row, col int) ContextMask {

	return (*c)[PixelPos{row, col}]

}

/// Factory of context table
func newContextMask(N int) *ContextMask {

	findAll := false
	distance := 0
	totalFound := 0

	cm := new(ContextMask)

	for {

		distance++
		found := 0
		secDistance := 0

		for secDistance <= distance && totalFound+found < N {

			if distance > secDistance && totalFound+found < N {

				*cm = append(*cm, PixelPos{0 - secDistance, 0 - distance})
				found++

			}
			if totalFound+found < N {

				*cm = append(*cm, PixelPos{0 - distance, 0 - secDistance})
				found++

			}
			if secDistance > 0 && totalFound+found < N {

				*cm = append(*cm, PixelPos{0 - distance, 0 + secDistance})
				found++

			}
			secDistance++

		}

		totalFound += found
		if totalFound == N && found == 0 {
			findAll = true
		}

		if findAll {
			break
		}
	}

	return cm

}

/// Apply mask to determine local context of PixelPos
func applyMask(p PixelPos, cm ContextMask, width, height int) *[]PixelPos {

	ret := new([]PixelPos)
	for _, dot := range cm {

		candidate := PixelPos{p.Row + dot.Row, p.Col + dot.Col}
		if candidate.Row >= 0 &&
			candidate.Row < height &&
			candidate.Col >= 0 &&
			candidate.Col < width {
			(*ret) = append(*ret, candidate)
		} else {
			continue
		}

	}
	return ret

}

func NewContextTable(N, width, height int) *ContextTable {

	table := make(ContextTable, height*width)
	mask := newContextMask(N)

	for row := 0; row < height; row++ {

		for col := 0; col < width; col++ {

			if _, present := table[PixelPos{row, col}]; !present {

				neighbours := applyMask(PixelPos{row, col}, *mask, width, height)
				table[PixelPos{row, col}] = *neighbours

			}

		}

	}

	return &table

}
