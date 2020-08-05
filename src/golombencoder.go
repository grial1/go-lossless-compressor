package src

import "math"

/// Rice mapping of the prediction error
func RiceMapping(e *GreyImage) *GreyImage {

	m := NewGreyImage(e.GetHeight(), e.GetWidth())

	for row := 0; row < e.GetHeight(); row++ {
		for col := 0; col < e.GetWidth(); col++ {
			if e.GetPixel(row, col) >= 0 {
				m.SetPixel(row, col, 2*e.GetPixel(row, col))
			} else {
				m.SetPixel(row, col, -2*e.GetPixel(row, col)-1)
			}
		}
	}

	return m

}

/// Golomb code order (m=1<<k)
func GetCodeOrder(p PixelPos, t *ContextTable, e *GreyImage) int {

	if len((*t)[p]) == 0 {
		return 1 << 3
	} else {

		/// Count number of prediction errors
		n := len((*t)[p])

		/// Sum absolute value of prediction errors
		A := 0

		for _, v := range (*t)[p] {

			A += int(math.Abs(float64(e.GetValueFromPixelPos(v))))

		}

		k := 0
		for (n << k) < A {
			k++
		}

		return 1 << k
	}

}
