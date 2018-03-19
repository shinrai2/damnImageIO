package bmp

import (
	"fmt"
	"os"

	head "../head"
	matrix "../matrix"
	util "../util"
)

// Read bmp data.
func Read(filePath *string) {
	fmt.Println("Input path: ", *filePath)
	f, err := os.Open(*filePath)
	util.Check(err)
	/* Load time. */
	bitmapFileHeader := head.BitmapFileHeader{
		BfType:      string(readNextBytes(f, 2)),
		BfSize:      util.ByteArr2int32(readNextBytes(f, 4)),
		BfReserved1: util.ByteArr2int16(readNextBytes(f, 2)),
		BfReserved2: util.ByteArr2int16(readNextBytes(f, 2)),
		BfOffBits:   util.ByteArr2int32(readNextBytes(f, 4)),
	}
	bmpInfoHeader := head.BmpInfoHeader{
		BiSize:          util.ByteArr2int32(readNextBytes(f, 4)),
		BiWidth:         util.ByteArr2int32(readNextBytes(f, 4)),
		BiHeight:        util.ByteArr2int32(readNextBytes(f, 4)),
		BiPlanes:        util.ByteArr2int16(readNextBytes(f, 2)),
		BiBitCount:      util.ByteArr2int16(readNextBytes(f, 2)),
		BiCompression:   util.ByteArr2int32(readNextBytes(f, 4)),
		BiSizeImage:     util.ByteArr2int32(readNextBytes(f, 4)),
		BiXPelsPerMeter: util.ByteArr2int32(readNextBytes(f, 4)),
		BiYPelsPerMeter: util.ByteArr2int32(readNextBytes(f, 4)),
		BiClrUsed:       util.ByteArr2int32(readNextBytes(f, 4)),
		BiClrImportant:  util.ByteArr2int32(readNextBytes(f, 4)),
	}
	rgbQuads := make([]head.RgbQuads, 0)
	if bmpInfoHeader.BiBitCount <= 8 { // Grayscale: <=8
		for i := 0; i < (1 << uint(bmpInfoHeader.BiBitCount)); i++ {
			rgbQuads = append(rgbQuads, head.RgbQuads{
				RgbBlue:     readNextBytes(f, 1)[0],
				RgbGreen:    readNextBytes(f, 1)[0],
				RgbRed:      readNextBytes(f, 1)[0],
				RgbReserved: readNextBytes(f, 1)[0],
			})
		}
	}
	dataSizePerLine := (int(bmpInfoHeader.BiWidth)*int(bmpInfoHeader.BiBitCount) + 31) / 8
	dataSizePerLine = (dataSizePerLine / 4) * 4 // Make sure the variable is a multiple of four.
	imageData := make([]head.ImageLine, 0, bmpInfoHeader.BiHeight)
	for i := 0; i < int(bmpInfoHeader.BiHeight); i++ { // Loop for read all pixel data.
		imageData = append(imageData, head.ImageLine{
			ImageByteArr: readNextBytes(f, dataSizePerLine),
		})
	}
	mdata := make([]uint8, 0, int(bmpInfoHeader.BiWidth)*int(bmpInfoHeader.BiHeight))
	for i := int(bmpInfoHeader.BiHeight) - 1; i >= 0; i-- {
		mdata = append(mdata, imageData[i].Format2intArr(int(bmpInfoHeader.BiBitCount), int(bmpInfoHeader.BiWidth))...)
	}
	imageData = nil
	mdimen := []int{int(bmpInfoHeader.BiWidth), int(bmpInfoHeader.BiHeight)}
	imgMatrix := matrix.Create(mdimen, mdata)
	/* In the end we get the structure what we want */
	bmpData := head.BmpData{
		BitmapFileHeader: bitmapFileHeader,
		BmpInfoHeader:    bmpInfoHeader,
		RgbQuads:         rgbQuads,
		ImgMatrix:        imgMatrix,
	}
	_ = bmpData
	/* Show time. */
	/* BITMAPFILEHEADER */
	fmt.Println("bfType:\t\t", bitmapFileHeader.BfType)
	fmt.Println("bfSize:\t\t", bitmapFileHeader.BfSize)
	fmt.Println("bfReserved1:\t", bitmapFileHeader.BfReserved1)
	fmt.Println("bfReserved2:\t", bitmapFileHeader.BfReserved2)
	fmt.Println("bfOffBits:\t", bitmapFileHeader.BfOffBits)
	fmt.Println("")
	/* BMP_INFOHEADER */
	fmt.Println("biSize:\t\t", bmpInfoHeader.BiSize)
	fmt.Println("biWidth:\t", bmpInfoHeader.BiWidth)
	fmt.Println("biHeight:\t", bmpInfoHeader.BiHeight)
	fmt.Println("biPlanes:\t", bmpInfoHeader.BiPlanes)
	fmt.Println("biBitCount:\t", bmpInfoHeader.BiBitCount)
	fmt.Println("biCompression:\t", bmpInfoHeader.BiCompression)
	fmt.Println("biSizeImage:\t", bmpInfoHeader.BiSizeImage)
	fmt.Println("biXPelsPerMeter:", bmpInfoHeader.BiXPelsPerMeter)
	fmt.Println("biYPelsPerMeter:", bmpInfoHeader.BiYPelsPerMeter)
	fmt.Println("biClrUsed:\t", bmpInfoHeader.BiClrUsed)
	fmt.Println("biClrImportant:\t", bmpInfoHeader.BiClrImportant)
	fmt.Println("")
	/* RGBQUAD */
	for i := 0; i < len(rgbQuads); i++ {
		fmt.Println("RGBQUAD ", i, ":\t", rgbQuads[i].Format())
	}
	fmt.Println("")
	/* IMAGEDATA */
	imgMatrix.Printx()
	f.Close() // avoid OOM
}

// readNextBytes read the next x bytes in 'os.File'.
func readNextBytes(f *os.File, size int) []byte {
	bx := make([]byte, size)
	_, err := f.Read(bx)
	util.Check(err)
	return bx
}
