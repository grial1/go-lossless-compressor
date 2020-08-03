// Package greyImage has the implementation of the type GreyImage, used to parse
// images in PGM file format.
package src

import (
	"io"
	"os"
	"reflect"
	"strconv"
)

// Grey scale image interface
type GreyImage struct {

	// Number of coloumns
	width int

	// Number of rows
	height int

	// Pixeles of the image
	imageMat []int16
}

/// Accessor
func (g *GreyImage) GetWidth() int { return g.width }

/// Accessor
func (g *GreyImage) GetHeight() int { return g.height }

/// Accessor
func (g *GreyImage) GetPixel(row, col int) int16 {
	return g.imageMat[row*(g.height-(g.height-g.width))+col]
}

// GreyImage factory function
func NewGreyImage(height, width int) *GreyImage {

	gc := GreyImage{width, height, make([]int16, height*width)}
	return &gc

}

// GreyImage fectory function using an image file
func NewGreyImageFromFile(filePath string) *GreyImage {

	f, err := os.Open(filePath)
	if err != nil {

		panic(err)

	}
	charBuff := make([]byte, 1)
	var size int = 1

	gotType := false
	gotWidth := false
	gotHeight := false
	gotMaxValue := false

	var height, width int

	var lineBuffer []byte
	var dataBuffer []int16 /// Where data is saved

	for size > 0 {

		size, err = f.Read(charBuff)
		if err != nil && err != io.EOF {

			panic(err)

		}

		if !gotType {

			if charBuff[0] == byte('\n') {

				/// ignore type (P5)
				gotType = !gotType
				if !reflect.DeepEqual(lineBuffer, []byte{'P', '5'}) {

					panic("The file does not have PGM raw -binary- format")

				}
				lineBuffer = nil
				continue

			} else {

				/// Add char to line
				lineBuffer = append(lineBuffer, charBuff[0])
				continue

			}

		} else if !gotWidth {

			if charBuff[0] == byte(' ') {

				gotWidth = !gotWidth
				width, err = strconv.Atoi(string(lineBuffer))
				if err != nil {
					panic(err)
				}
				lineBuffer = nil
				continue

			} else {

				/// Add char to line
				lineBuffer = append(lineBuffer, charBuff[0])
				continue

			}

		} else if !gotHeight {

			if charBuff[0] == byte('\n') {

				gotHeight = !gotHeight
				height, err = strconv.Atoi(string(lineBuffer))
				if err != nil {
					panic(err)
				}
				lineBuffer = nil
				continue

			} else {

				/// Add char to line
				lineBuffer = append(lineBuffer, charBuff[0])
				continue

			}

		} else if !gotMaxValue {

			if charBuff[0] == byte('\n') {

				/// Ignore max value
				gotMaxValue = !gotMaxValue
				if !reflect.DeepEqual(lineBuffer, []byte("255")) {

					panic("The file does not have PGM raw -binary- format")

				}
				lineBuffer = nil
				continue

			} else {

				/// Add char to line
				lineBuffer = append(lineBuffer, charBuff[0])
				continue

			}

		} else {

			/// Add data
			dataBuffer = append(dataBuffer, int16(charBuff[0]))
			continue

		}

	}

	f.Close()
	gc := GreyImage{width, height, dataBuffer}

	return &gc
}

/// Function used to save GreyImage to a file
func (g *GreyImage) Save(filename string) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {

		panic(err)

	}

	/// Write file type
	f.Write([]byte{'P', '5', '\n'})

	/// Write image size
	f.Write([]byte(strconv.Itoa(g.width)))
	f.Write([]byte(" "))
	f.Write([]byte(strconv.Itoa(g.height)))
	f.Write([]byte("\n"))

	/// Write max value
	f.Write([]byte("255"))
	f.Write([]byte("\n"))

	/// Write image data
	var dataBuffer []byte
	for _, px := range g.imageMat {

		dataBuffer = append(dataBuffer, byte(px))

	}
	f.Write(dataBuffer)
	f.Close()

}
