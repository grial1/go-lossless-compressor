// Package that calls the compression function from src and
// provides a simple, text-based, user interface
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/grial1/go-lossless-compressor/src"
)

func main() {

	args := os.Args
	if len(args) < 3 {
		fmt.Println("USAGE: ./compressor.bin <image-path> <N>")
		fmt.Println("Decompress with the script ubuntu_compressor.bash (See docs)")
		os.Exit(-1)
	}

	fmt.Printf("Compressing %s\n", args[1])
	N, _ := strconv.Atoi(args[2])
	basename, fType := filepath.Base(args[1]), "P5"
	ext := strings.Index(basename, ".")
	if ext > -1 {

		if basename[ext+1:] == "ppm" {
			fType = "P6"
		}
		basename = basename[:ext]

	}

	if fType == "P5" {

		gc := src.NewGreyImageFromFile(args[1])
		src.Compress(gc, basename, "P5", gc.GetWidth(), gc.GetHeight(), N, src.Red)

	} else {

		cm := src.NewColourImageFromFile(args[1])
		src.Compress(cm.GetRedImage(), basename, "P6", cm.GetWidth(), cm.GetHeight(), N, src.Red)
		src.Compress(cm.GetGreenImage(), basename, "P6", cm.GetWidth(), cm.GetHeight(), N, src.Green)
		src.Compress(cm.GetBlueImage(), basename, "P6", cm.GetWidth(), cm.GetHeight(), N, src.Blue)

	}

	fmt.Println("Operation finished")

}
