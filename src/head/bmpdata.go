package head

import (
	"errors"
	"fmt"
	"os"

	util "../util"
	"gonum.org/v1/gonum/mat"
)

// BitmapFileHeader record the file information.
type BitmapFileHeader struct {
	BfType      string
	BfSize      uint32
	BfReserved1 uint16
	BfReserved2 uint16
	BfOffBits   uint32
}

// BmpInfoHeader record the image information.
type BmpInfoHeader struct {
	BiSize          uint32
	BiWidth         int32 // LONG, has signed bit.
	BiHeight        int32 // LONG
	BiPlanes        uint16
	BiBitCount      uint16 // ignore 16 and 32
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32 // LONG
	BiYPelsPerMeter int32 // LONG
	BiClrUsed       uint32
	BiClrImportant  uint32
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

// ReadPixelData read all pixel data and transform to a dense array(3-dimen).
func ReadPixelData(f1 *os.File, width, height int32, biBitCount uint16) []*mat.Dense {
	var count int
	lineLength := util.GetLengthOfLine(width, biBitCount) // the actually length of line byte in bmp file.
	dimz := int((biBitCount-1)/8 + 1)                     // calculate the dimension z
	r := make([][]float64, dimz)                          // the data before packing
	rd := make([]*mat.Dense, dimz)                        // the actual return value
	for iz := 0; iz < dimz; iz++ {
		r[iz] = make([]float64, int(width*height))
	}
	for ih := 0; ih < int(height); ih++ { // loop
		vl := util.ReadNextBytes(f1, lineLength) // read a line data.
		for _, v := range vl {                   // loop for each byte of line
			switch int(biBitCount) {
			// use RGBQUAD
			case 1: // one pixel one bit
				for i := 7; i >= 0; i-- {
					if count >= int(width)*(ih+1) { // Key code, skip the remaining blank parts of line.
						break
					}
					y := (int(height)-(count/int(width))-1)*int(width) + count%int(width) // Calculate the actual dimension two
					r[0][y] = float64((v & (byte(1) << uint(i))) >> uint(i))
					count++
				}
			case 4: // one pixel four bits
				for i := 1; i >= 0; i-- {
					if count >= int(width)*(ih+1) {
						break
					}
					y := (int(height)-(count/int(width))-1)*int(width) + count%int(width)
					r[0][y] = float64((v & (byte(15) << uint(i*4))) >> uint(i*4))
					count++
				}
			case 8: // one pixel eight bits
				if count >= int(width)*(ih+1) {
					break
				}
				y := (int(height)-(count/int(width))-1)*int(width) + count%int(width)
				r[0][y] = float64(v)
				count++
			// not use RGBQUAD
			case 16: // three pixels sixteen bits
				util.Check(errors.New("Unsupported 16-bits type now"))
			case 24: // three pixels twenty-four bits
				if count >= int(width)*(ih+1)*3 {
					break
				}
				y := (int(height)-(count/3/int(width))-1)*int(width) + count/3%int(width)
				r[count%3][y] = float64(v)
				count++
			case 32: // four pixels thirty-two bits
				if count >= int(width)*(ih+1)*4 {
					break
				}
				y := (int(height)-(count/4/int(width))-1)*int(width) + count/4%int(width)
				r[count%4][y] = float64(v)
				count++
			default:
				util.Check(errors.New("Unsupported biBitCount type"))
			}
		}
	}
	// Packing
	for id, vd := range r {
		rd[id] = mat.NewDense(int(height), int(width), vd)
	}
	return rd
}
