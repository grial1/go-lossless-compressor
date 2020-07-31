// Package greyImage has the implementation of the type GreyImage, used to parse
// images in PGM file format.
package src

// Grey scale image interface
type GreyImage struct {

	// Number of coloumns
	width int

	// Number of rows
	height int

	// Pixeles of the image
	imageMat []int16
}
