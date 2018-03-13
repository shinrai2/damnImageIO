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

// Create a new matrix after the rules check.
func Create(dimension []int, data []int) Matrix {
	l, err := util.MultiplyByEach(dimension)
	util.Check(err)
	if l != len(data) {
		panic(errors.New("The data and dimensions of the matrix do not match"))
	}
	return Matrix{
		dimension,
		data,
	}
}

// Sum each element of matrix and return the result.
func (matrix Matrix) Sum() int {
	var r int
	for i := 0; i < len(matrix.data); i++ {
		r = r + matrix.data[i]
	}
	return r
}

// Add add the corresponding elements of two matrices.
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
