package main

import (
	"fmt"

	"github.com/grial1/go-lossless-compressor/src"
)

func main() {

	gc := src.NewGreyImageFromFile("test.pgm")

	fmt.Printf("%d, %d\n", gc.GetWidth(), gc.GetHeight())

	fmt.Printf("%d, %d\n", gc.GetPixel(0, 0), gc.GetPixel(40, 100))

	gc.Save("saved_test.pgm")

}
