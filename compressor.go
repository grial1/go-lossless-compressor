package main

import (
	"fmt"

	"github.com/grial1/go-lossless-compressor/src"
)

func main() {

	gc := src.NewGreyImage(4, 4)
	gc.SetPixel(0, 0, 30)
	gc.SetPixel(0, 1, 130)
	gc.SetPixel(0, 2, 200)
	gc.SetPixel(0, 3, 250)
	gc.SetPixel(1, 0, 150)
	gc.SetPixel(1, 1, 10)
	gc.SetPixel(1, 2, 210)
	gc.SetPixel(1, 3, 35)
	gc.SetPixel(2, 0, 130)
	gc.SetPixel(2, 1, 230)
	gc.SetPixel(2, 2, 20)
	gc.SetPixel(2, 3, 50)
	gc.SetPixel(3, 0, 60)
	gc.SetPixel(3, 1, 63)
	gc.SetPixel(3, 2, 60)
	gc.SetPixel(3, 3, 5)
	//gc.Save("minitest.pgm")

	t := src.NewContextTable(5, gc.GetWidth(), gc.GetHeight())

	for k, v := range *t {

		fmt.Printf("central px: %d,%d: ", k.Row, k.Col)
		for _, n := range v {

			fmt.Printf("\n\t px: %d,%d", n.Row, n.Col)

		}
		fmt.Println()

	}

	//gc := src.NewGreyImageFromFile("test.pgm")
	//pred := src.FixedPrediction(*gc)
	//error := gc.Sub(*pred)

	//fmt.Printf("Dimensions %d, %d\n", gc.GetWidth(), gc.GetHeight())

	//error.Save("prediction_error.pgm")

}
