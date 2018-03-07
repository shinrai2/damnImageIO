package main

import (
	"flag"
	"fmt"
	"os"

	head "./head"
	util "./util"
)

// main func.
func main() {
	filePath := flag.String("path", "source/t_1.bmp", "the path of file.")
	flag.Parse()
	fmt.Println("Input path: ", *filePath)
	f, err := os.Open(*filePath)
	util.Check(err)
	// read BITMAPFILEHEADER values and put them values to the structure.
	BitmapFileHeader := head.BITMAPFILEHEADER{
		string(util.ReadNextBytes(f, 2)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int16(util.ReadNextBytes(f, 2)),
		util.ByteArr2int16(util.ReadNextBytes(f, 2)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
	}
	BmpInfoHeader := head.BMP_INFOHEADER{
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int16(util.ReadNextBytes(f, 2)),
		util.ByteArr2int16(util.ReadNextBytes(f, 2)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
	}
	RgbQuads := make([]head.RGBQUAD, 0)
	if BmpInfoHeader.BiBitCount <= 8 { // Grayscale: <=8
		for i := 0; i < (1 << uint(BmpInfoHeader.BiBitCount)); i++ {
			Rgbq := head.RGBQUAD{
				int8(util.ReadNextBytes(f, 1)[0]),
				int8(util.ReadNextBytes(f, 1)[0]),
				int8(util.ReadNextBytes(f, 1)[0]),
				int8(util.ReadNextBytes(f, 1)[0]),
			}
			RgbQuads = append(RgbQuads, Rgbq)
		}
	}
	/* BITMAPFILEHEADER */
	fmt.Println("bfType:\t\t", BitmapFileHeader.BfType)
	fmt.Println("bfSize:\t\t", BitmapFileHeader.BfSize)
	fmt.Println("bfReserved1:\t", BitmapFileHeader.BfReserved1)
	fmt.Println("bfReserved2:\t", BitmapFileHeader.BfReserved2)
	fmt.Println("bfOffBits:\t", BitmapFileHeader.BfOffBits)
	fmt.Println("")
	/* BMP_INFOHEADER */
	fmt.Println("biSize:\t\t", BmpInfoHeader.BiSize)
	fmt.Println("biWidth:\t", BmpInfoHeader.BiWidth)
	fmt.Println("biHeight:\t", BmpInfoHeader.BiHeight)
	fmt.Println("biPlanes:\t", BmpInfoHeader.BiPlanes)
	fmt.Println("biBitCount:\t", BmpInfoHeader.BiBitCount)
	fmt.Println("biCompression:\t", BmpInfoHeader.BiCompression)
	fmt.Println("biSizeImage:\t", BmpInfoHeader.BiSizeImage)
	fmt.Println("biXPelsPerMeter:", BmpInfoHeader.BiXPelsPerMeter)
	fmt.Println("biYPelsPerMeter:", BmpInfoHeader.BiYPelsPerMeter)
	fmt.Println("biClrUsed:\t", BmpInfoHeader.BiClrUsed)
	fmt.Println("biClrImportant:\t", BmpInfoHeader.BiClrImportant)
	fmt.Println("")
	/* RGBQUAD */
	for i := 0; i < len(RgbQuads); i++ {
		fmt.Println("RGBQUAD ", i, ":\t(", RgbQuads[i].RgbBlue, ",",
			RgbQuads[i].RgbGreen, ",", RgbQuads[i].RgbRed, ",", RgbQuads[i].RgbReserved, ")")
	}
	fmt.Println("")

	f.Close()
}
