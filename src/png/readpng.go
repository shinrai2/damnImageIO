package png

import (
	"fmt"
	"os"

	head "../head"
	util "../util"
)

// Read png data.
func Read(filePath *string) {
	fmt.Println("Input path: ", *filePath)
	f, err := os.Open(*filePath)
	util.Check(err)
	/* Load time. */
	pngFileSignature := head.PngFileSignature(util.ReadNextBytes(f, 8)) // begining of file.
	// IHDR Chunk
	length := util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	chunkTypeCode := string(util.ReadNextBytes(f, 4))
	iHDR := head.IHDR{
		Width:       util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
		Height:      util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
		BitDepth:    uint8(util.ReadNextBytes(f, 1)[0]),
		ColorType:   uint8(util.ReadNextBytes(f, 1)[0]),
		Compression: util.ReadNextBytes(f, 1)[0],
		Filter:      util.ReadNextBytes(f, 1)[0],
		Interlace:   util.ReadNextBytes(f, 1)[0],
	}
	crc := util.ReadNextBytes(f, 4)
	iHDRChunk := head.Chunk{
		Length:        length,
		ChunkTypeCode: chunkTypeCode,
		ChunkData:     iHDR,
		CRC:           crc,
	}
	// sRGB Chunk
	length = util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	chunkTypeCode = string(util.ReadNextBytes(f, 4))
	sRGB := head.SRGB{RenderingIntent: util.ReadNextBytes(f, 1)[0]}
	crc = util.ReadNextBytes(f, 4)
	sRGBChunk := head.Chunk{
		Length:        length,
		ChunkTypeCode: chunkTypeCode,
		ChunkData:     sRGB,
		CRC:           crc,
	}
	// gAMA Chunk
	length = util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	chunkTypeCode = string(util.ReadNextBytes(f, 4))
	gAMA := head.GAMA{
		Gamma: util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
	}
	crc = util.ReadNextBytes(f, 4)
	gAMAChunk := head.Chunk{
		Length:        length,
		ChunkTypeCode: chunkTypeCode,
		ChunkData:     gAMA,
		CRC:           crc,
	}
	// pHYs Chunk
	length = util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	chunkTypeCode = string(util.ReadNextBytes(f, 4))
	pHYs := head.PHYS{
		Xaxis:         util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
		Yaxis:         util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
		UnitSpecifier: util.ReadNextBytes(f, 1)[0],
	}
	crc = util.ReadNextBytes(f, 4)
	pHYsChunk := head.Chunk{
		Length:        length,
		ChunkTypeCode: chunkTypeCode,
		ChunkData:     pHYs,
		CRC:           crc,
	}
	// IDAT Chunk
	length = util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	chunkTypeCode = string(util.ReadNextBytes(f, 4))
	iDAT := util.ReadNextBytes(f, int(length))
	crc = util.ReadNextBytes(f, 4)
	iDATChunk := head.Chunk{
		Length:        length,
		ChunkTypeCode: chunkTypeCode,
		ChunkData:     iDAT,
		CRC:           crc,
	}
	// IEND
	length = util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	chunkTypeCode = string(util.ReadNextBytes(f, 4))
	crc = util.ReadNextBytes(f, 4)
	iENDChunk := head.Chunk{
		Length:        length,
		ChunkTypeCode: chunkTypeCode,
		ChunkData:     nil,
		CRC:           crc,
	}
	/* Show time. */
	fmt.Println("PngFileSignature:\t", pngFileSignature)

	fmt.Println("iHDRLength:\t\t", iHDRChunk.Length)
	fmt.Println("iHDRChunkTypeCode:\t", iHDRChunk.ChunkTypeCode)
	fmt.Println("iHDRWidth:\t\t", iHDRChunk.ChunkData.(head.IHDR).Width)
	fmt.Println("iHDRHeight:\t\t", iHDRChunk.ChunkData.(head.IHDR).Height)
	fmt.Println("iHDRBitDepth:\t\t", iHDRChunk.ChunkData.(head.IHDR).BitDepth)
	fmt.Println("iHDRColorType:\t\t", iHDRChunk.ChunkData.(head.IHDR).ColorType)
	fmt.Println("iHDRCompression:\t", iHDRChunk.ChunkData.(head.IHDR).Compression)
	fmt.Println("iHDRFilter:\t\t", iHDRChunk.ChunkData.(head.IHDR).Filter)
	fmt.Println("iHDRInterlace:\t\t", iHDRChunk.ChunkData.(head.IHDR).Interlace)
	fmt.Println("iHDRCRC:\t\t", iHDRChunk.CRC)

	fmt.Println("\nsRGBLength:\t\t", sRGBChunk.Length)
	fmt.Println("sRGBChunkTypeCode:\t", sRGBChunk.ChunkTypeCode)
	fmt.Println("RenderingIntent:\t", sRGBChunk.ChunkData.(head.SRGB).RenderingIntent)
	fmt.Println("sRGBCRC:\t\t", sRGBChunk.CRC)

	fmt.Println("\ngAMALength:\t\t", gAMAChunk.Length)
	fmt.Println("gAMAChunkTypeCode:\t", gAMAChunk.ChunkTypeCode)
	fmt.Println("gAMAGamma:\t\t", gAMAChunk.ChunkData.(head.GAMA).Gamma)
	fmt.Println("gAMACRC:\t\t", gAMAChunk.CRC)

	fmt.Println("\npHYsLength:\t\t", pHYsChunk.Length)
	fmt.Println("pHYsChunkTypeCode:\t", pHYsChunk.ChunkTypeCode)
	fmt.Println("pHYsXaxis:\t\t", pHYsChunk.ChunkData.(head.PHYS).Xaxis)
	fmt.Println("pHYsYaxis:\t\t", pHYsChunk.ChunkData.(head.PHYS).Yaxis)
	fmt.Println("pHYsUnitSpecifier:\t", pHYsChunk.ChunkData.(head.PHYS).UnitSpecifier)
	fmt.Println("pHYsCRC:\t\t", pHYsChunk.CRC)

	fmt.Println("\niDATLength:\t\t", iDATChunk.Length)
	fmt.Println("iDATChunkTypeCode:\t", iDATChunk.ChunkTypeCode)
	fmt.Println("iDATData:\t\t", iDATChunk.ChunkData.([]byte))
	fmt.Println("iDATCRC:\t\t", iDATChunk.CRC)

	fmt.Println("\niENDLength:\t\t", iENDChunk.Length)
	fmt.Println("iENDChunkTypeCode:\t", iENDChunk.ChunkTypeCode)
	fmt.Println("iENDCRC:\t\t", iENDChunk.CRC)
}
