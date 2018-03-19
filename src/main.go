package main

import (
	"flag"
	"fmt"

	bmp "./bmp"
	util "./util"
)

// main func.
func main() {
	filePath := flag.String("path", "source/white.bmp", "the path of file.")
	flag.Parse()
	bmp.Read(filePath)

	var a int32 = 57662
	ar := util.Int2ByteArr(a)
	fmt.Println(ar)
}
