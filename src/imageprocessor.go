package src

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Compress(g *GreyImage, filename, imageType string, width, height, N int /*, rgb Colour, transform bool*/) {

	/// Open compress file to write
	var f *os.File

	if imageType == "P5" /* || rgb == RED */ {

		nf, err := os.OpenFile(filename+".loco", os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		f = nf
		f.Write([]byte(fmt.Sprintf("%s\n", imageType)))
		f.Write([]byte(fmt.Sprintf("%d %d\n", width, height)))
		f.Write([]byte(fmt.Sprintf("%d\n", N)))
		/*if imageType == "P6" {
			// TODO
		}*/

	} /* else {
		// TODO: Append others colours to the file
	} */

	/// 1. Apply MED
	pred := FixedPrediction(*g)

	/// 2. Find prediction error
	gError := g.Sub(*pred)

	/// 3. Find rice mapping
	rm := RiceMapping(gError)

	/// 4. Calculate context table for a given N size
	table := NewContextTable(N, width, height)

	/// 5. Calculate the compressed image code
	var b string
	for row := 0; row < g.GetHeight(); row++ {
		for col := 0; col < g.GetWidth(); col++ {

			pxVal := PixelPos{row, col}
			m := GetCodeOrder(pxVal, table, gError)
			var rem int = int(rm.GetValueFromPixelPos(pxVal)) % m
			var quot int = int(rm.GetValueFromPixelPos(pxVal)) / m

			/// Compressed code
			var cod string = ""

			/// bin(rice mod m)
			if m > 1 {

				k := 0
				rest := 0
				cod = "0"
				if rem > 0 {
					k = int(math.Floor(math.Log2(float64(rem))))
					rest = rem - (1 << k)
					cod = "1"
				}

				for i := 0; i < k; i++ {
					cod += "0"
				}

				for rest > 0 {
					k = int(math.Floor(math.Log2(float64(rest))))
					rest -= 1 << k
					cod = cod[:len(cod)-1-k] + "1" + cod[len(cod)-k:]
				}

				pending := int(math.Ceil(math.Log2(float64(m)))) - len(cod)
				for j := 0; j < pending; j++ {
					cod = "0" + cod
				}

			}

			/// uni(rice div m)
			for i := 0; i < quot; i++ {
				cod += "0"
			}
			cod += "1"

			b += cod

			if len(b) > 7 {

				i := 0
				var data []byte
				for ; i < len(b)/8; i++ {

					v, err := strconv.ParseInt(b[8*i:8*(i+1)], 2, 0)
					if err != nil {
						panic(err)
					}

					data = append(data, byte(v))

				}
				f.Write(data)
				b = b[8*i:]

			}
		}
	}

	/// Padding
	if len(b) > 0 {

		for len(b)%8 != 0 {
			b += "0"
		}

		var data []byte
		for i := 0; i < len(b)/8; i++ {

			v, err := strconv.ParseInt(b[8*i:8*(i+1)], 2, 0)
			if err != nil {
				panic(err)
			}

			data = append(data, byte(v))

		}
		f.Write(data)

	}
	f.Close()

}
