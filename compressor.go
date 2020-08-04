package main

import (
	"fmt"

	"github.com/grial1/go-lossless-compressor/src"
)

func main() {

	gc := src.NewGreyImageFromFile("test.pgm")
	pred := src.FixedPrediction(*gc)
	error := gc.Sub(*pred)

	fmt.Printf("Dimensions %d, %d\n", gc.GetWidth(), gc.GetHeight())

	error.Save("prediction_error.pgm")

}
