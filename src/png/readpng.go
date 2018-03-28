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
	iHDRLength := util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	iHDRChunkTypeCode := string(util.ReadNextBytes(f, 4))
	iHDR := head.IHDR{
		Width:       util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
		Height:      util.ByteArr2int32u(util.ReadNextBytes(f, 4), false),
		BitDepth:    uint8(util.ReadNextBytes(f, 1)[0]),
		ColorType:   uint8(util.ReadNextBytes(f, 1)[0]),
		Compression: util.ReadNextBytes(f, 1)[0],
		Filter:      util.ReadNextBytes(f, 1)[0],
		Interlace:   util.ReadNextBytes(f, 1)[0],
	}
	iHDRCRC := util.ByteArr2int32u(util.ReadNextBytes(f, 4), false)
	iHDRChunk := head.Chunk{
		Length:        iHDRLength,
		ChunkTypeCode: iHDRChunkTypeCode,
		ChunkData:     iHDR,
		CRC:           iHDRCRC,
	}
	//
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
}
