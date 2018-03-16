package head

import (
	matrix "../matrix"
)

// BmpData records all bmp data include the head data and the pixel data.
type BmpData struct {
	bitmapFileHeader BitmapFileHeader
	bmpInfoHeader    BmpInfoHeader
	rgbQuads         RgbQuads
	data             matrix.Matrix
}
