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

		if totalFound == N && found == 0 {
			findAll = true
		}

		if findAll {
			break
		}
	}

	return cm

}

func applyMask(p PixelPos, cm ContextMask, width, height int) *[]PixelPos {
	// TODO
}
