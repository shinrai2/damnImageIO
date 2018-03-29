package util

import (
	"errors"
	"os"
)

// Check panic if the func returns the error value is not nil.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// ByteArr2int32 transform byte Array to int32, ignore the sign bit(!).
func ByteArr2int32(byteArr []byte, direction bool) int32 {
	var r int32
	for i, v := range byteArr {
		if direction == true {
			r += int32(v) << uint(i*8)
		} else {
			r = r << 8
			r += int32(v)
		}
	}
	return r
}

// ByteArr2int16 transform byte Array to int16, ignore the sign bit(!).
func ByteArr2int16(byteArr []byte, direction bool) int16 {
	var r int16
	for i, v := range byteArr {
		if direction == true {
			r += int16(v) << uint(i*8)
		} else {
			r = r << 8
			r += int16(v)
		}
	}
	return r
}

// ByteArr2int32u transform byte Array to !!uint32, ignore the sign bit(!).
func ByteArr2int32u(byteArr []byte, direction bool) uint32 {
	var r uint32
	for i, v := range byteArr {
		if direction == true {
			r += uint32(v) << uint(i*8)
		} else {
			r = r << 8
			r += uint32(v)
		}
	}
	return r
}

// ByteArr2int16u transform byte Array to !!uint16, ignore the sign bit(!).
func ByteArr2int16u(byteArr []byte, direction bool) uint16 {
	var r uint16
	for i, v := range byteArr {
		if direction == true {
			r += uint16(v) << uint(i*8)
		} else {
			r = r << 8
			r += uint16(v)
		}
	}
	return r
}

// ReadNextBytes read the next x bytes in 'os.File'.
func ReadNextBytes(f1 *os.File, size int) []byte {
	bx := make([]byte, size)
	_, err := f1.Read(bx)
	Check(err)
	return bx
}

// Int2ByteArray transform int to byte Array.
func Int2ByteArray(intt interface{}, direction bool) []byte {
	var ix int
	var vx uint
	var r []byte
	switch intt.(type) {
	default:
		Check(errors.New("unexpected type"))
	case int16:
		ix = 2
		vx = uint(intt.(int16))
	case uint16:
		ix = 2
		vx = uint(intt.(uint16))
	case int32:
		ix = 4
		vx = uint(intt.(int32))
	case uint32:
		ix = 4
		vx = uint(intt.(uint32))
	}
	r = make([]byte, 0, ix)
	for i := 0; i < ix; i++ {
		if direction == true {
			r = append(r, byte(vx>>uint(i*8))) // bmp
		} else {
			r = append(r, byte(vx>>uint((ix-1-i)*8))) // png
		}
	}
	return r
}

// CompareIntArr Compare two arrays.
func CompareIntArr(a1 []int, a2 []int) (bool, string) {
	r := true    // default is true.
	var d string // default is nil.
	if len(a1) != len(a2) {
		r = false
		d = "The number of dimensions of the matrix is not the same"
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			r = false
			d = "The length of some dimensions in the matrix is not the same"
			break
		}
	}
	return r, d
}

// MultiplyByEach check for all element and multiply them to be the return value.
func MultiplyByEach(arr []int) (int, error) {
	r := 1
	var err error
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			err = errors.New("Zero exist")
			break
		}
		r = arr[i] * r
	}
	return r, err
}

// GetLengthOfLine ..
func GetLengthOfLine(biWidth int32, biBitCount uint16) int {
	dataSizePerLine := (int(biWidth)*int(biBitCount) + 31) / 8
	return (dataSizePerLine / 4) * 4 // Make sure the variable is a multiple of four.
}
