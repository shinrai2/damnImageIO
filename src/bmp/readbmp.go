package bmp

import (
	"fmt"
	"os"

	head "../head"
	util "../util"
)

// Read bmp data.
func Read(filePath *string) {
	fmt.Println("Input path: ", *filePath)
	f, err := os.Open(*filePath)
	util.Check(err)
	/* Load time. */
	bitmapFileHeader := head.BitmapFileHeader{
		BfType:      string(util.ReadNextBytes(f, 2)),
		BfSize:      util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
		BfReserved1: util.ByteArr2int16u(util.ReadNextBytes(f, 2), true),
		BfReserved2: util.ByteArr2int16u(util.ReadNextBytes(f, 2), true),
		BfOffBits:   util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
	}
	bmpInfoHeader := head.BmpInfoHeader{
		BiSize:          util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
		BiWidth:         util.ByteArr2int32(util.ReadNextBytes(f, 4), true), // LONG
		BiHeight:        util.ByteArr2int32(util.ReadNextBytes(f, 4), true), // LONG
		BiPlanes:        util.ByteArr2int16u(util.ReadNextBytes(f, 2), true),
		BiBitCount:      util.ByteArr2int16u(util.ReadNextBytes(f, 2), true),
		BiCompression:   util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
		BiSizeImage:     util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
		BiXPelsPerMeter: util.ByteArr2int32(util.ReadNextBytes(f, 4), true), // LONG
		BiYPelsPerMeter: util.ByteArr2int32(util.ReadNextBytes(f, 4), true), // LONG
		BiClrUsed:       util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
		BiClrImportant:  util.ByteArr2int32u(util.ReadNextBytes(f, 4), true),
	}
	rgbQuads := make([]head.RgbQuads, 0)
	if bmpInfoHeader.BiBitCount <= 8 { // Grayscale: <=8
		for i := 0; i < (1 << uint(bmpInfoHeader.BiBitCount)); i++ {
			rgbQuads = append(rgbQuads, head.RgbQuads{
				RgbBlue:     util.ReadNextBytes(f, 1)[0],
				RgbGreen:    util.ReadNextBytes(f, 1)[0],
				RgbRed:      util.ReadNextBytes(f, 1)[0],
				RgbReserved: util.ReadNextBytes(f, 1)[0],
			})
		}
	}
	pixelDense := head.ReadPixelData(f, bmpInfoHeader.BiWidth, bmpInfoHeader.BiHeight, bmpInfoHeader.BiBitCount)
	/* In the end we get the structure what we want */
	bmpData := head.BmpData{
		BitmapFileHeader: bitmapFileHeader,
		BmpInfoHeader:    bmpInfoHeader,
		RgbQuads:         rgbQuads,
		ImgDense:         pixelDense,
	}
	/* Show time. */
	/* BITMAPFILEHEADER */
	fmt.Println("bfType:\t\t", bmpData.BitmapFileHeader.BfType)
	fmt.Println("bfSize:\t\t", bmpData.BitmapFileHeader.BfSize)
	fmt.Println("bfReserved1:\t", bmpData.BitmapFileHeader.BfReserved1)
	fmt.Println("bfReserved2:\t", bmpData.BitmapFileHeader.BfReserved2)
	fmt.Println("bfOffBits:\t", bmpData.BitmapFileHeader.BfOffBits)
	fmt.Println("")
	/* BMP_INFOHEADER */
	fmt.Println("biSize:\t\t", bmpData.BmpInfoHeader.BiSize)
	fmt.Println("biWidth:\t", bmpData.BmpInfoHeader.BiWidth)
	fmt.Println("biHeight:\t", bmpData.BmpInfoHeader.BiHeight)
	fmt.Println("biPlanes:\t", bmpData.BmpInfoHeader.BiPlanes)
	fmt.Println("biBitCount:\t", bmpData.BmpInfoHeader.BiBitCount)
	fmt.Println("biCompression:\t", bmpData.BmpInfoHeader.BiCompression)
	fmt.Println("biSizeImage:\t", bmpData.BmpInfoHeader.BiSizeImage)
	fmt.Println("biXPelsPerMeter:", bmpData.BmpInfoHeader.BiXPelsPerMeter)
	fmt.Println("biYPelsPerMeter:", bmpData.BmpInfoHeader.BiYPelsPerMeter)
	fmt.Println("biClrUsed:\t", bmpData.BmpInfoHeader.BiClrUsed)
	fmt.Println("biClrImportant:\t", bmpData.BmpInfoHeader.BiClrImportant)
	fmt.Println("")
	/* RGBQUAD */
	for i, v := range bmpData.RgbQuads {
		fmt.Println("RGBQUAD ", i, ":\t", v.Format())
	}
	fmt.Println("")
	/* IMAGEDATA */
	for _, vid := range bmpData.ImgDense {
		fmt.Println(*vid)
	}
	f.Close()
}
