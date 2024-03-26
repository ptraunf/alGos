package main

import "testing"

func TestNewTable(t *testing.T) {
	h := 5
	w := 4
	table := newTable(h, w, 0)
	if table == nil {
		t.Fatal("Table is nil")
	}
	actualLen := len(table.data)
	expectedLen := w * h
	if expectedLen != actualLen {
		t.Fatalf("\nExpected %v elements; Got %v elements.\n", expectedLen, actualLen)
	}
	// t.Logf("Table:\n%v\n", table)
}

func TestTableSet(t *testing.T) {
	h := 5
	w := 4
	initial := -1
	table := newTable(h, w, initial)
	expected := 5
	table.set(2, 3, expected)
	actual := table.at(2, 3)
	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t%v\n", expected, actual)
	}
}
func TestTableAt(t *testing.T) {
	h := 5
	w := 4
	initial := -1
	table := newTable(h, w, initial)

	actual := table.at(2, 3)
	if initial != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t%v\n", initial, actual)
	}
}

func TestEditDistance(t *testing.T) {
	a := "spongebob"
	b := "spingborb"
	expected := 3
	actual := EditDistance(a, b)
	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
