package bmp

import (
	"bufio"
	"os"

	head "../head"
	util "../util"
)

// writeNextBytes ..
func writeNextBytes(wr *bufio.Writer, pdata []byte) {
	_, err := wr.Write(pdata)
	util.Check(err)
	wr.Flush()
}

// Write ..
func Write(filename string, data head.BmpData) {
	f, err := os.Create(filename)
	util.Check(err)
	wb := bufio.NewWriter(f)
	_ = wb

	f.Close()
}
