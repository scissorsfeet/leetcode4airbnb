package _28__Longest_Consecutive_Sequence

import "testing"

func TestA(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	if 4 != longestConsecutive(nums) {
		t.Fail()
	}
}

func TestB(t *testing.T) {
	nums := []int{100, 4, 5, 1, 2, 3}
	if 5 != longestConsecutive(nums) {
		t.Fail()
	}
}

func TestC(t *testing.T) {
	nums := []int{100, 400, 500, 200, 300}
	if 1 != longestConsecutive(nums) {
		t.Fail()
	}
}

func TestD(t *testing.T) {
	nums := []int{-1, 9, -3, -6, 7, -8, -6, 2, 9, 2, 3, -2, 4, -1, 0, 6, 1, -9, 6, 8, 6, 5, 2}
	if 13 != longestConsecutive(nums) {
		t.Fail()
	}
}

func TestE(t *testing.T) {
	nums := []int{0, 1, 0, 0}
	if 2 != longestConsecutive(nums) {
		t.Fail()
	}
}
