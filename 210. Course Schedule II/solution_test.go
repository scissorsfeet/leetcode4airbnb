package _10__Course_Schedule_II

import "testing"

func TestA(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {3, 4}, {2, 5}, {4, 2}}
	t.Log(findOrder(6, prerequisites))
}

func TestB(t *testing.T) {
	prerequisites := [][]int{{1, 0}}
	t.Log(findOrder(2, prerequisites))
}

func TestC(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {0, 1}}
	t.Log(findOrder(2, prerequisites))
}
