package _07__Course_Schedule

import "testing"

func TestA(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {3, 4}, {2, 5}, {4, 2}}
	if !canFinish(6, prerequisites) {
		t.Fatal()
	}
}

func TestB(t *testing.T) {
	prerequisites := [][]int{{1, 0}}
	if !canFinish(2, prerequisites) {
		t.Fatal()
	}
}

func TestC(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {0, 1}}
	if canFinish(2, prerequisites) {
		t.Fatal()
	}
}
