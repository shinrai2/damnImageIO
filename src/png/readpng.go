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
	iHDRCRC := util.ReadNextBytes(f, 4)
	iHDRChunk := head.Chunk{
		Length:        iHDRLength,
		ChunkTypeCode: iHDRChunkTypeCode,
		ChunkData:     iHDR,
		CRC:           iHDRCRC,
	}
	_ = pngFileSignature
	_ = iHDRChunk
}
