package main

import (
	"fmt"
	"math"
)

type table[T any] struct {
	height int
	width  int
	data   []T
}

func newTable[T any](h int, w int, initial T) *table[T] {
	// data := []T{}
	data := make([]T, h*w)
	for i := 0; i < h*w; i++ {
		data[i] = initial
	}
	return &table[T]{
		height: h,
		width:  w,
		data:   data,
	}
}

func (t *table[T]) at(row int, col int) T {
	return t.data[t.width*row+col]
}

func (t *table[T]) set(row int, col int, value T) {
	t.data[t.width*row+col] = value
}

func (t *table[T]) String() string {
	var s string
	for i := 0; i < t.height*t.width; i++ {
		val := t.data[i]
		s = s + fmt.Sprintf("\t%v", val)
		if i%t.width == (t.width - 1) {
			s = s + "\n"
		}
	}
	return s
}

func EditDistance(a string, b string) int {
	// # a b c d e f
	// a 0 1 2 3 4 5
	// z 1 1 2 3 4 5
	// c 2 2 1 2 3 4
	// e 3 3 2 2 2 3
	// d 4 4 3 2 3 3 <- Final Result

	// # # a b c d e f
	// # 0 1 2 3 4 5 6
	// a 1 0 1 2 3 4 5
	// z 2 1 1 2 3 4 5
	// c 3 2 2 1 2 3 4
	// e 4 3 3 2 2 2 3
	// d 5 4 4 3 2 3 3 <- Final Result

	// | diag | top
	// |------+-----
	// | left | current <- diag OR 1 + min(diag, left, top)
	rows := len(a) + 1
	cols := len(b) + 1
	as := []rune(a)
	bs := []rune(b)
	t := newTable(rows, cols, 0)
	// initialize table - first row and col
	for r := 0; r < rows; r++ {
		t.set(r, 0, r)
	}
	for c := 0; c < cols; c++ {
		t.set(0, c, c)
	}
	i := 1
	j := 1
	for i < rows && j < cols {
		if as[i] == bs[j] {
			// Gets the diagonal, or 0 if in ocrner
			t.set(i, j, t.at(i-1, j-1))
			i++
			j++
		} else {
			minNeighbor := int(
				math.Min(
					math.Min(float64(t.at(i-1, j-1)), float64(t.at(i-1, j))),
					float64(t.at(i, j-1)),
				),
			)
			t.set(i, j, minNeighbor)
		}
	}
	return 0
}
