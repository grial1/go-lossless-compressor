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
