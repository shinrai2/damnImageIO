package head

import (
	"fmt"
)

// BitmapFileHeader record the file information.
type BitmapFileHeader struct {
	BfType      string
	BfSize      int32
	BfReserved1 int16
	BfReserved2 int16
	BfOffBits   int32
}

// BmpInfoHeader record the image information.
type BmpInfoHeader struct {
	BiSize          int32
	BiWidth         int32 // LONG, has signed bit.
	BiHeight        int32 // LONG
	BiPlanes        int16
	BiBitCount      int16
	BiCompression   int32
	BiSizeImage     int32
	BiXPelsPerMeter int32 // LONG
	BiYPelsPerMeter int32 // LONG
	BiClrUsed       int32
	BiClrImportant  int32
}

// RgbQuads record the RGB information.
type RgbQuads struct {
	RgbBlue     int8
	RgbGreen    int8
	RgbRed      int8
	RgbReserved int8
}

// Format the RgbQuads like (0, 0, 0, 0)
func (rgbQuads RgbQuads) Format() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", rgbQuads.RgbBlue,
		rgbQuads.RgbGreen, rgbQuads.RgbRed, rgbQuads.RgbReserved)
}
