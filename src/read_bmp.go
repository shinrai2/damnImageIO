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
	/* Load time. */
	bitmapFileHeader := head.BitmapFileHeader{
		string(util.ReadNextBytes(f, 2)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
		util.ByteArr2int16(util.ReadNextBytes(f, 2)),
		util.ByteArr2int16(util.ReadNextBytes(f, 2)),
		util.ByteArr2int32(util.ReadNextBytes(f, 4)),
	}
	bmpInfoHeader := head.BmpInfoHeader{
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
	rgbQuads := make([]head.RgbQuads, 0)
	if bmpInfoHeader.BiBitCount <= 8 { // Grayscale: <=8
		for i := 0; i < (1 << uint(bmpInfoHeader.BiBitCount)); i++ {
			rgbQuads = append(rgbQuads, head.RgbQuads{
				int8(util.ReadNextBytes(f, 1)[0]),
				int8(util.ReadNextBytes(f, 1)[0]),
				int8(util.ReadNextBytes(f, 1)[0]),
				int8(util.ReadNextBytes(f, 1)[0]),
			})
		}
	}
	dataSizePerLine := (int(bmpInfoHeader.BiWidth)*int(bmpInfoHeader.BiBitCount) + 31) / 8
	dataSizePerLine = (dataSizePerLine / 4) * 4 // Make sure the variable is a multiple of four.
	imageData := make([]head.ImageLine, 0, bmpInfoHeader.BiHeight)
	for i := 0; i < int(bmpInfoHeader.BiHeight); i++ { // Loop for read all pixel data.
		imageData = append(imageData, head.ImageLine{
			util.ReadNextBytes(f, dataSizePerLine),
		})
	}
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
	sizeOfData := len(imageData)
	fmt.Println("len of imageData is: ", sizeOfData)
	oneLine := imageData[sizeOfData-1].Format(
		int(bmpInfoHeader.BiBitCount), int(bmpInfoHeader.BiWidth))
	fmt.Println("the first line of data is: ", oneLine)
	fmt.Println("the len of first line of data is: ", len(oneLine))
	f.Close()
}
