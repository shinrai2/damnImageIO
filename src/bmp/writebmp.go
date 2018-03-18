package bmp

import (
	"bufio"
	"errors"
	"os"

	util "../util"
)

type BmpFileWriteBuffer struct {
	buffer *bufio.Writer
}

// CreateFile ..
func CreateFile(filename string) BmpFileWriteBuffer {
	var writeBuffer BmpFileWriteBuffer
	if util.CheckFileIsExist(filename) == false {
		f, err := os.Create(filename)
		defer f.Close()
		util.Check(err)
		writeBuffer = BmpFileWriteBuffer{buffer: bufio.NewWriter(f)}
	} else {
		util.Check(errors.New("File already exists"))
	}
	return writeBuffer
}
