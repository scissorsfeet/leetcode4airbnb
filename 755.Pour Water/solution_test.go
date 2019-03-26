package _55_Pour_Water

import "testing"

func TestA(t *testing.T) {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	K := 3
	V := 4
	pourWater(heights, K, V)
}
