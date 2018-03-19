package head

import (
	matrix "../matrix"
)

// BmpData records all bmp data include the head data and the pixel data.
type BmpData struct {
	BitmapFileHeader BitmapFileHeader
	BmpInfoHeader    BmpInfoHeader
	RgbQuads         []RgbQuads
	ImgMatrix        matrix.Matrix
}
