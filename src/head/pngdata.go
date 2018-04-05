package head

// PngFileSignature is the first eight bytes of a PNG file.
type PngFileSignature []byte

// Chunk is the basic structure
type Chunk struct {
	Length        uint32 // 4 bytes
	ChunkTypeCode string // 4 bytes
	ChunkData     interface{}
	CRC           []byte
}

// IHDR structure
type IHDR struct {
	Width       uint32
	Height      uint32
	BitDepth    uint8
	ColorType   uint8
	Compression byte
	Filter      byte
	Interlace   byte
}

// PLTE structure
type PLTE struct {
	Red   byte
	Green byte
	Blue  byte
}

// SRGB structure
type SRGB struct {
	RenderingIntent byte
}

// GAMA structure
type GAMA struct {
	Gamma uint32
}

// PHYS structure
type PHYS struct {
	Xaxis         uint32
	Yaxis         uint32
	UnitSpecifier byte
}

// IDAT structure
type IDAT struct {
}
