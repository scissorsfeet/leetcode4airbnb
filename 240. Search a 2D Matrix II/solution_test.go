package _40__Search_a_2D_Matrix_II

import "testing"

func TestA(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if !searchMatrix(matrix, 5) {
		t.Fatal()
	}
}

func TestB(t *testing.T) {
	matrix := [][]int{
		{-1, 3},
	}
	if !searchMatrix(matrix, 3) {
		t.Fatal()
	}
}

func TestC(t *testing.T) {
	matrix := [][]int{
		{-1, 3},
	}
	if searchMatrix(matrix, -3) {
		t.Fatal()
	}
}

func TestD(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if searchMatrix(matrix, 0) {
		t.Fatal()
	}
}

func TestE(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if searchMatrix(matrix, 100) {
		t.Fatal()
	}
}
