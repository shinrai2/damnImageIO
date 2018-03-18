package util

import (
	"bufio"
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
func ByteArr2int32(byteArr []byte) int32 {
	var r int32
	for i, v := range byteArr {
		r += int32(v) << uint(i*8)
	}
	return int32(r)
}

// ByteArr2int16 transform byte Array to int16, ignore the sign bit(!).
func ByteArr2int16(byteArr []byte) int16 {
	var r int16
	for i, v := range byteArr {
		r += int16(v) << uint(i*8)
	}
	return int16(r)
}

// ReadNextBytes read the next x bytes in 'os.File'.
func ReadNextBytes(br *bufio.Reader, size int) []byte {
	bx := make([]byte, size)
	_, err := br.Read(bx)
	Check(err)
	return bx
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

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
