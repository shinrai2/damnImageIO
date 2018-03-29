package head

// PngFileSignature is the first eight bytes of a PNG file.
type PngFileSignature []byte

type Chunk struct {
	Length        uint32 // 4 bytes
	ChunkTypeCode string // 4 bytes
	ChunkData     interface{}
	CRC           []byte
}

type IHDR struct {
	Width       uint32
	Height      uint32
	BitDepth    uint8
	ColorType   uint8
	Compression byte
	Filter      byte
	Interlace   byte
}

type PLTE struct {
	Red   byte
	Green byte
	Blue  byte
}

// IEND is the last twelve bytes of a PNG file.
type IEND []byte
