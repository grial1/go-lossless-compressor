package src

type Image interface {
	GetHeight() int
	GetWidth() int
	//Add(i Image) *Image
	//Sub(i Image) *Image
	Save(filename string)
}
