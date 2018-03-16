package matrix

import (
	"testing"
)

func TestCheckA(t *testing.T) {
	m1 := Matrix{[]int{1, 2},
		[]int{1, 2}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	_, _, err := checkA(m1, m2)
	if err == nil {
		t.Error("Wrong calculation was accepted")
	}
}

func TestSum1(t *testing.T) {
	s := Matrix{[]int{2, 2},
		[]int{0, 0, 0, 0}}.Sum()
	if s != 0 {
		t.Error("The sum of zero matrix is not zero")
	}
}

func TestSum2(t *testing.T) {
	s := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}.Sum()
	if s != 10 {
		t.Error("The sum of matrix [1, 2, 3, 4] is not ten")
	}
}

func TestEqual1(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	if Equal(m1, m2) == true {
		t.Error("Wrong result of the Matrix Equal Func")
	}
}

func TestEqual2(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	if Equal(m1, m2) == false {
		t.Error("Wrong result of the Matrix Equal Func")
	}
}

func TestAdd(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	m3 := m1.Add(m2)
	if Equal(m3, Matrix{[]int{2, 2}, []int{6, 8, 10, 12}}) == false {
		t.Error("Wrong result of the Matrix Add Func")
	}
}

func TestSubtract(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	m3 := m2.Subtract(m1)
	if Equal(m3, Matrix{[]int{2, 2}, []int{4, 4, 4, 4}}) == false {
		t.Error("Wrong result of the Matrix Subtract Func")
	}
}

func TestMultiply(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	m3 := m2.Multiply(m1)
	if Equal(m3, Matrix{[]int{2, 2}, []int{5, 12, 21, 32}}) == false {
		t.Error("Wrong result of the Matrix Multiply Func")
	}
}

func TestDivide(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	m3 := m2.Divide(m1)
	if Equal(m3, Matrix{[]int{2, 2}, []int{5, 3, 2, 2}}) == false {
		t.Error("Wrong result of the Matrix Divide Func")
	}
}
