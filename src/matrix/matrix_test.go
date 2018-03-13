package matrix

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	m1 := Matrix{[]int{2, 2},
		[]int{1, 2, 3, 4}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	_, err := m1.Add(m2)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAdd2(t *testing.T) {
	m1 := Matrix{[]int{1, 2},
		[]int{1, 2}}
	m2 := Matrix{[]int{2, 2},
		[]int{5, 6, 7, 8}}
	_, err := m1.Add(m2)
	if err != nil {
		t.Error(err.Error())
	}
}
