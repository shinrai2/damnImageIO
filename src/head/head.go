package head

type BITMAPFILEHEADER struct {
	BfType      string
	BfSize      int32
	BfReserved1 int16
	BfReserved2 int16
	BfOffBits   int32
}
type BMP_INFOHEADER struct {
	BiSize          int32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        int16
	BiBitCount      int16
	BiCompression   int32
	BiSizeImage     int32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       int32
	BiClrImportant  int32
}
type RGBQUAD struct {
	RgbBlue     int8
	RgbGreen    int8
	RgbRed      int8
	RgbReserved int8
}
