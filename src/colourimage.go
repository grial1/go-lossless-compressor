package src

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

// RGB colour enum
type Colour uint8

const (
	Red Colour = iota
	Green
	Blue
)

// Colour image object
// formed by three planes
type ColourImage struct {
	pRed,
	pGreen,
	pBlue *GreyImage
}

/// Accessor
func (c *ColourImage) GetWidth() int { return c.pRed.GetWidth() }

/// Accessor
func (c *ColourImage) GetHeight() int { return c.pRed.GetHeight() }

/// Accessor
func (c *ColourImage) GetRedPixel(row, col int) int16 {
	if row < c.GetHeight() && col < c.GetWidth() {
		return c.pRed.GetPixel(row, col)
	} else {
		msg := fmt.Sprintf("Error: Index out of bound: RED row: %d, col: %d", row, col)
		panic(msg)
	}
}

/// Accessor
func (c *ColourImage) GetGreenPixel(row, col int) int16 {
	if row < c.GetHeight() && col < c.GetWidth() {
		return c.pGreen.GetPixel(row, col)
	} else {
		msg := fmt.Sprintf("Error: Index out of bound: GREEN row: %d, col: %d", row, col)
		panic(msg)
	}
}

/// Accessor
func (c *ColourImage) GetBluePixel(row, col int) int16 {
	if row < c.GetHeight() && col < c.GetWidth() {
		return c.pBlue.GetPixel(row, col)
	} else {
		msg := fmt.Sprintf("Error: Index out of bound: BLUE row: %d, col: %d", row, col)
		panic(msg)
	}
}

/// Mutators
func (c *ColourImage) SetRedPixel(row, col int, value int16) {
	c.pRed.SetPixel(row, col, value)
}

/// Mutators
func (c *ColourImage) SetGreenPixel(row, col int, value int16) {
	c.pGreen.SetPixel(row, col, value)
}

/// Mutators
func (c *ColourImage) SetBluePixel(row, col int, value int16) {
	c.pBlue.SetPixel(row, col, value)
}

// ColourImage factory function
func NewColourImage(height, width int) *ColourImage {

	pRed := NewGreyImage(height, width)
	pGreen := NewGreyImage(height, width)
	pBlue := NewGreyImage(height, width)

	c := ColourImage{pRed, pGreen, pBlue}

	return &c

}

// ColourImage factory function using an image file
func NewColourImageFromFile(filePath string) *ColourImage {

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

	var height, width, colour int = 0, 0, int(Red)

	var lineBuffer []byte
	var redDataBuffer, greenDataBuffer, blueDataBuffer []int16

	for size > 0 {

		size, err = f.Read(charBuff)
		if err != nil && err != io.EOF {

			panic(err)

		}

		if !gotType {

			if charBuff[0] == byte('\n') {

				/// ignore type (P5)
				gotType = !gotType
				if !reflect.DeepEqual(lineBuffer, []byte{'P', '6'}) {

					panic("The file does not have PPM raw -binary- format")

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
			switch Colour(colour) {
			case Red:
				redDataBuffer = append(redDataBuffer, int16(charBuff[0]))
				colour++
			case Green:
				greenDataBuffer = append(greenDataBuffer, int16(charBuff[0]))
				colour++
			case Blue:
				blueDataBuffer = append(blueDataBuffer, int16(charBuff[0]))
				colour = int(Red)
			}
			continue

		}

	}

	f.Close()
	fmt.Printf("%d %d %d \n", len(redDataBuffer), len(greenDataBuffer), len(blueDataBuffer))
	pRed := GreyImage{width, height, redDataBuffer}
	pGreen := GreyImage{width, height, greenDataBuffer}
	pBlue := GreyImage{width, height, blueDataBuffer}
	cm := ColourImage{&pRed, &pGreen, &pBlue}

	return &cm

}

// Save ColourImage to a PPM file
func (c *ColourImage) Save(filename string) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {

		panic(err)

	}

	/// Write file type
	f.Write([]byte{'P', '6', '\n'})

	/// Write image size
	f.Write([]byte(strconv.Itoa(c.GetWidth())))
	f.Write([]byte(" "))
	f.Write([]byte(strconv.Itoa(c.GetHeight())))
	f.Write([]byte("\n"))

	/// Write max value
	f.Write([]byte("255"))
	f.Write([]byte("\n"))

	/// Write image data
	var dataBuffer []byte
	for row := 0; row < c.GetHeight(); row++ {
		for col := 0; col < c.GetWidth(); col++ {
			for colour := 0; colour < 3; colour++ {

				switch Colour(colour % 3) {
				case Red:
					dataBuffer = append(dataBuffer, byte(c.GetRedPixel(row, col)))
				case Green:
					dataBuffer = append(dataBuffer, byte(c.GetGreenPixel(row, col)))
				case Blue:
					dataBuffer = append(dataBuffer, byte(c.GetBluePixel(row, col)))
				}

			}
		}
	}

	f.Write(dataBuffer)
	f.Close()

}
