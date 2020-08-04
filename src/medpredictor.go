package src

/// MED predition for border detection
func FixedPrediction(g GreyImage) *GreyImage {

	p := NewGreyImage(g.GetHeight(), g.GetWidth())

	for row := 0; row < g.GetHeight(); row++ {

		for col := 0; col < g.GetWidth(); col++ {

			var a, b, c int16

			/// Initialize values
			if row == 0 {
				b = 0
				c = 0
			} else {
				b = g.GetPixel(row-1, col)
				if col == 0 {
					if row-2 > 0 {
						c = g.GetPixel(row-2, col)
					} else {
						c = 0
					}
				} else {
					c = g.GetPixel(row-1, col-1)
				}
			}
			if col == 0 {
				a = b
			} else {
				a = g.GetPixel(row, col-1)
			}
			/// Make predictions
			if c >= a && c >= b {

				if a > b {
					p.SetPixel(row, col, b)
				} else {
					p.SetPixel(row, col, a)
				}

			} else if c <= a && c <= b {

				if a > b {
					p.SetPixel(row, col, a)
				} else {
					p.SetPixel(row, col, b)
				}

			} else {

				p.SetPixel(row, col, a+b-c)

			}
		}

	}

	return p

}

/// func UpdateFixedPrediction() //TODO
