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
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

// Format the RgbQuads like (0, 0, 0, 0)
func (rgbQuads RgbQuads) Format() string {
	return fmt.Sprint("(", rgbQuads.RgbBlue, ", ", rgbQuads.RgbGreen,
		", ", rgbQuads.RgbRed, ", ", rgbQuads.RgbReserved, ")")
}

// ImageLine one Line of Pixel, may have some padding at the end.
type ImageLine struct {
	ImageByteArr []byte
}

// Format the ImageLine
func (imageLine ImageLine) Format(biBitCount int, biWidth int) ([]byte, error) {
	var layer []byte
	var err error
	var count int
	if biBitCount > 8 { // Non-grayscale
		layer = make([]byte, 0, biWidth*3)
	} else { // Grayscale
		layer = make([]byte, 0, biWidth)
	}
	for _, v := range imageLine.ImageByteArr {
		switch biBitCount {
		case 1:
			for i := 7; i >= 0; i-- {
				if count < biWidth {
					layer = append(layer, (v&(byte(1)<<uint(i)))>>uint(i))
					count++
				}
			}
		case 4:
			for i := 1; i >= 0; i-- {
				if count < biWidth {
					layer = append(layer, (v&(byte(15)<<uint(i*4)))>>uint(i*4))
					count++
				}
			}
		case 8:
			if count < biWidth {
				layer = append(layer, v)
				count++
			}
		// case 16:
		// case 24:
		// case 32:
		default:
			err = errors.New("Unsupported biBitCount type")
		}
	}
	return layer, err
}
