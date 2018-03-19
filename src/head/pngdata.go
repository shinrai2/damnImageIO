package head

// PngFileSignature is the first eight bytes of a PNG file.
type PngFileSignature [8]byte

type IHDR struct {
	Width       uint32
	Height      uint32
	BitDepth    uint8
	ColorType   uint8
	Compression byte
	Filter      byte
	Interlace   byte
}

type IEND [12]byte
