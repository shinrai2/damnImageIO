package png

import (
	"fmt"
	"os"

	util "../util"
)

func Read(filePath *string) {
	fmt.Println("Input path: ", *filePath)
	f, err := os.Open(*filePath)
	util.Check(err)
	/* Load time. */
}
