package main

import "fmt"

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
	return 0
}
