package png

import (
	"fmt"
	"os"

	head "../head"
	util "../util"
)

func Read(filePath *string) {
	fmt.Println("Input path: ", *filePath)
	f, err := os.Open(*filePath)
	util.Check(err)
	/* Load time. */
	pngFileSignature := head.PngFileSignature(util.ReadNextBytes(f, 8)) // begining of file.
	// wid :=
	// iHDR := head.Chunk{}

}
