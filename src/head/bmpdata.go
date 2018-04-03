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

// // ImageLine one Line of Pixel, may have some padding at the end.
// type ImageLine struct {
// 	ImageByteArr []byte
// }

// // Format the ImageLine to a one dimension byte matrix.
// func (imageLine ImageLine) Format(biBitCount int, biWidth int) []byte {
// 	var layer []byte
// 	var count int
// 	if biBitCount == 32 { // Alpha
// 		layer = make([]uint8, biWidth*4)
// 	} else if biBitCount > 8 { // Non-grayscale
// 		layer = make([]uint8, biWidth*3)
// 	} else { // Grayscale
// 		layer = make([]uint8, biWidth)
// 	}
// 	for _, v := range imageLine.ImageByteArr {
// 		switch biBitCount {
// 		// use RGBQUAD
// 		case 1:
// 			for i := 7; i >= 0; i-- {
// 				if count < biWidth {
// 					layer[count] = uint8((v & (byte(1) << uint(i))) >> uint(i))
// 					// layer = append(layer, )
// 					count++
// 				}
// 			}
// 		case 4:
// 			for i := 1; i >= 0; i-- {
// 				if count < biWidth {
// 					layer[count] = uint8((v & (byte(15) << uint(i*4))) >> uint(i*4))
// 					// layer = append(layer, )
// 					count++
// 				}
// 			}
// 		case 8:
// 			if count < biWidth {
// 				layer[count] = uint8(v)
// 				// layer = append(layer, )
// 				count++
// 			}
// 		// not use RGBQUAD
// 		case 16:
// 			if count < biWidth {
// 		case 24:
// 		case 32:
// 		default:
// 			util.Check(errors.New("Unsupported biBitCount type"))
// 		}
// 	}
// 	return layer
// }
func ReadPixelData(f1 *os.File, width, height int32, biBitCount uint16) []*mat.Dense {
	var count int
	lineLength := util.GetLengthOfLine(width, biBitCount) // the actually length of line byte in bmp file.
	dimz := int((biBitCount-1)/8 + 1)                     // calculate the dimension z
	r := make([][]float64, dimz)
	rd := make([]*mat.Dense, dimz)
	for iz := 0; iz < dimz; iz++ {
		r[iz] = make([]float64, int(width*height))
	}
	for ih := 0; ih < int(height); ih++ { // loop
		vl := util.ReadNextBytes(f1, lineLength) // read a line data.
		for _, v := range vl {                   // loop for each byte of line
			switch int(biBitCount) {
			// use RGBQUAD
			case 1:
				for i := 7; i >= 0; i-- {
					if count >= int(width)*(ih+1) { // Key code, skip the remaining blank parts of line.
						break
					}
					y := (int(height)-(count/int(width))-1)*int(width) + count%int(width)
					r[0][y] = float64((v & (byte(1) << uint(i))) >> uint(i))
					count++
				}
			case 4:
				for i := 1; i >= 0; i-- {
					if count >= int(width)*(ih+1) {
						break
					}
					y := (int(height)-(count/int(width))-1)*int(width) + count%int(width)
					r[0][y] = float64((v & (byte(15) << uint(i*4))) >> uint(i*4))
					count++
				}
			case 8:
				if count >= int(width)*(ih+1) {
					break
				}
				y := (int(height)-(count/int(width))-1)*int(width) + count%int(width)
				r[0][y] = float64(v)
				count++
			// not use RGBQUAD
			case 16:
			case 24:
			case 32:
			default:
				util.Check(errors.New("Unsupported biBitCount type"))
			}
		}
	}
	for id, vd := range r {
		rd[id] = mat.NewDense(int(height), int(width), vd)
	}
	return rd
}
