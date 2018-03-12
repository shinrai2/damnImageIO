package matrix

import (
	"errors"

	util "../util"
)

// Matrix structure to be determined...
type Matrix struct {
	dimension []int
	data      []int
}

func (matrixA Matrix) Add(matrixB Matrix) (Matrix, error) {
	var c Matrix
	var err error
	lena := len(matrixA.dimension)
	lend := len(matrixA.data)
	if lena == 0 || len(matrixB.dimension) == 0 {
		err = errors.New("Empty matrix(s) exist")
	} else {
		r, d := util.CompareIntArr(matrixA.dimension, matrixB.dimension)
		if r == false {
			err = errors.New(d)
		} else {
			dimen := make([]int, lena)
			data := make([]int, lend)
			copy(dimen, matrixA.dimension)
			for i := 0; i < lend; i++ {
				data = append(data, matrixA.data[i]+matrixB.data[i])
			}
			c = Matrix{
				dimen,
				data,
			}
		}
	}
	return c, err
}

func (matrixA Matrix) Subtract(matrixB Matrix) (Matrix, error) {
	var c Matrix
	return c, nil
}

func (matrixA Matrix) Multiply(matrixB Matrix) (Matrix, error) {
	var c Matrix
	return c, nil
}

func (matrixA Matrix) Divide(matrixB Matrix) (Matrix, error) {
	var c Matrix
	return c, nil
}
