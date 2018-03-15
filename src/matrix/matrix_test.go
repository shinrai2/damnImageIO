package matrix

import (
	"testing"
)

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

func TestCheckA(t *testing.T) {
	m1 := Matrix{[]int{1, 2},
		[]int{1, 2}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	err, _, _ := checkA(m1, m2)
	if err == nil {
		t.Error("Wrong calculation was accepted")
	}
}

func TestAdd(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	m3 := m1.Add(m2)
	if m3.Sum() != 36 {
		t.Error("Wrong result of the Matrix Add Func.")
	}
}
