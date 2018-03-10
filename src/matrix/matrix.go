package matrix

// piece like one dimension matrix.
type piece struct {
	width int
	data  []interface{}
}

// Matrix contains some pieces for making up a matrix structure.
type Matrix struct {
	dimension []int
	data      []piece
}
