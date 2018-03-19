package matrix

import (
	"errors"
	"fmt"

	util "../util"
)

// Matrix structure to be determined...
type Matrix struct {
	dimension []int
	data      []uint8
}

// Create a new matrix after the rules check.
func Create(dimension []int, data []uint8) Matrix {
	l, err := util.MultiplyByEach(dimension)
	util.Check(err)
	if l != len(data) {
		panic(errors.New("The data and dimensions of the matrix do not match"))
	}
	return Matrix{
		dimension: dimension,
		data:      data,
	}
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

// Printx ..
func (matrix Matrix) Printx() {
	if len(matrix.dimension) == 2 {
		for i := 0; i < len(matrix.data); i++ {
			if i%matrix.dimension[0] == 0 {
				fmt.Println("") // next line.
			}
			fmt.Print(matrix.data[i], " ")
		}
	} else if len(matrix.dimension) == 3 {
		for i := 0; i < len(matrix.data); i++ {
			xy := matrix.dimension[0] * matrix.dimension[1]
			if i%matrix.dimension[0] == 0 {
				fmt.Println("") // next line.
				if i%xy == 0 {
					fmt.Printf("[::%d]\n", i/xy)
				}
			}
			fmt.Print(matrix.data[i], " ")
		}
	}
}

// Sum each element of matrix and return the result.
func (matrix Matrix) Sum() uint8 {
	var r uint8
	for i := 0; i < len(matrix.data); i++ {
		r = r + matrix.data[i]
	}
	return r
}

// Equal or not between two matrices.
func Equal(matrixA Matrix, matrixB Matrix) bool {
	b := true
	_, lend, err := checkA(matrixA, matrixB)
	util.Check(err)
	for i := 0; i < lend; i++ {
		if matrixA.data[i] != matrixB.data[i] {
			b = false
			break
		}
	}
	return b
}

// Add the corresponding elements of two matrices.
func (matrix Matrix) Add(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]uint8, 0, lend) // don't forget this zero, otherwise it will lead to data length error.
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

// Subtract the corresponding elements of two matrices.
func (matrix Matrix) Subtract(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]uint8, 0, lend)
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

// Multiply the corresponding elements of two matrices.
func (matrix Matrix) Multiply(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]uint8, 0, lend)
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

// Divide the corresponding elements of two matrices.
func (matrix Matrix) Divide(matrixB Matrix) Matrix {
	var c Matrix
	lena, lend, err := checkA(matrix, matrixB)
	util.Check(err)
	dimen := make([]int, lena)
	data := make([]uint8, 0, lend)
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
