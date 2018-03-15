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

// checkA is a function that check if the two matrices have the same shape.
func checkA(matrix Matrix, matrixB Matrix) (int, int, error) {
	var err error
	lena := len(matrix.dimension)
	lend := len(matrix.data)
	if lena == 0 || len(matrixB.dimension) == 0 {
		err = errors.New("Empty matrix(s) exist")
	} else {
		r, d := util.CompareIntArr(matrix.dimension, matrixB.dimension)
		if r == false {
			err = errors.New(d)
		}
	}
	return lena, lend, err
}

// Add add the corresponding elements of two matrices.
func (matrix Matrix) Add(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]int, lend)
	copy(dimen, matrix.dimension)
	for i := 0; i < lend; i++ {
		data = append(data, matrix.data[i]+matrixB.data[i])
	}
	c = Matrix{
		dimen,
		data,
	}
	return c
}

func (matrix Matrix) Subtract(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]int, lend)
	copy(dimen, matrix.dimension)
	for i := 0; i < lend; i++ {
		data = append(data, matrix.data[i]-matrixB.data[i])
	}
	c = Matrix{
		dimen,
		data,
	}
	return c
}

func (matrix Matrix) Multiply(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]int, lend)
	copy(dimen, matrix.dimension)
	for i := 0; i < lend; i++ {
		data = append(data, matrix.data[i]*matrixB.data[i])
	}
	c = Matrix{
		dimen,
		data,
	}
	return c
}

func (matrix Matrix) Divide(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]int, lend)
	copy(dimen, matrix.dimension)
	for i := 0; i < lend; i++ {
		data = append(data, matrix.data[i]/matrixB.data[i])
	}
	c = Matrix{
		dimen,
		data,
	}
	return c
}
