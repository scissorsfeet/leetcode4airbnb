package _73__Sliding_Puzzle

import "testing"

func TestA(t *testing.T) {
	res := slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}})
	if res != 1 {
		t.Fatal(res)
	}
}

func TestB(t *testing.T) {
	res := slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}})
	if res != -1 {
		t.Fatal(res)
	}
}

func TestC(t *testing.T) {
	res := slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}})
	if res != 5 {
		t.Fatal(res)
	}
}
