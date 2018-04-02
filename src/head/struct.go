package head

import (
	"gonum.org/v1/gonum/mat"
)

// BmpData records all bmp data include the head data and the pixel data.
type BmpData struct {
	BitmapFileHeader BitmapFileHeader
	BmpInfoHeader    BmpInfoHeader
	RgbQuads         []RgbQuads
	ImgDense         []*mat.Dense
}

// PngData ..
type PngData struct {
}
