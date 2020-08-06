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
	gc := src.NewGreyImageFromFile(args[1])
	N, _ := strconv.Atoi(args[2])
	basename := filepath.Base(args[1])
	ext := strings.Index(basename, ".")
	if ext > -1 {
		basename = basename[:ext]
	}
	src.Compress(gc, basename, "P5", gc.GetWidth(), gc.GetHeight(), N /*, rgb Colour, transform bool*/)
	fmt.Println("Operation finished")

}
