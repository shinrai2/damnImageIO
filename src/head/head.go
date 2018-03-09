package head

import (
	"errors"
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
	BiBitCount      int16 // ignore 16 and 32
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

// ImageLine one Line of Pixel, may have some padding at the end.
type ImageLine struct {
	ImageByteArr []byte
}

// Format the ImageLine
func (imageLine ImageLine) Format(biBitCount int, biWidth int) ([]int8, int64) {
	var layer []int8
	var count int64
	if biBitCount > 8 { // Non-grayscale
		layer = make([]int8, 0, biWidth*3)
	} else { // Grayscale
		layer = make([]int8, 0, biWidth)
	}
	for _, v := range imageLine.ImageByteArr {
		switch biBitCount {
		case 1:
			for i := uint(7); i >= uint(0); i-- {
				// layer = append(layer, ((int8(v) & (1 << i)) >> i))
				_ = v
				count++
			}
		// case 4:
		// case 8:
		// case 24:
		default:
			panic(errors.New("Unsupported biBitCount type"))
		}
	}
	return layer, count
}
