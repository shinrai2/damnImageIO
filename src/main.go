package main

import (
	"flag"

	bmp "./bmp"
)

// main func.
func main() {
	filePath := flag.String("path", "source/t_1.bmp", "the path of file.")
	flag.Parse()
	bmp.Read(filePath)
}
