package main

import (
	"os"
	"testing"

	"github.com/grial1/go-lossless-compressor/src"
)

func TestCompress(t *testing.T) {

	src.Compress(src.NewGreyImageFromFile("test.pgm"), "grey_test", "P5", 512, 512, 5, src.Red)
	_, err := os.Stat("grey_test.loco")

	if err != nil {
		t.Error("Error when compressing test.pgm")
	}

	cm := src.NewColourImageFromFile("test.ppm")
	src.Compress(cm.GetRedImage(), "colour_test", "P6", 768, 512, 5, src.Red)
	src.Compress(cm.GetGreenImage(), "colour_test", "P6", 768, 512, 5, src.Green)
	src.Compress(cm.GetBlueImage(), "colour_test", "P6", 768, 512, 5, src.Blue)
	_, err = os.Stat("colour_test.loco")

	if err != nil {
		t.Error("Error when compressing test.pgm")
	}
}
