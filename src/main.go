package main

import (
	"flag"

	bmp "./bmp"
	png "./png"
)

// main func.
func main() {
	filePath := flag.String("path", "source/t_1.bmp", "the path of file.")
	flag.Parse()
	// Automatic selection method according to file suffix
	switch (*filePath)[len(*filePath)-3:] {
	case "bmp":
		bmp.Read(filePath)
	case "jpg":
	case "png":
		png.Read(filePath)
	}
}
