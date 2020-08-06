package main

import (
	"os"
	"testing"

	"github.com/grial1/go-lossless-compressor/src"
)

func TestCompress(t *testing.T) {

	src.Compress(src.NewGreyImageFromFile("test.pgm"), "test", "P5", 512, 512, 5)
	_, err := os.Stat("test.loco")

	if err != nil {
		t.Error("Error when compressing test.pgm")
	}

}
