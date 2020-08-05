package src

import (
	"bytes"
	"fmt"
	"os"
)

func compress(g *GreyImage, filename, imageType string, width, height, N int /*, rgb Colour, transform bool*/) {

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
	var b bytes.Buffer
	for row := 0; row < e.GetHeight(); row++ {
		for col := 0; col < e.GetWidth(); col++ {

			pxVal := PixelPos{row, col}
			m := GetCodeOrder(pxVal, table, gError)
			var rem int = int(rm.GetValueFromPixelPos(pxVal)) / m
			var quot int = int(rm.GetValueFromPixelPos(pxVal)) % m

			/// TODO: finish compression

		}
	}

}
