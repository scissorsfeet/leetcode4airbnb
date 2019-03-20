package _35__Asteroid_Collision

import "testing"

func TestA(t *testing.T) {
	nums := []int{5, 10, -5}
	t.Log(asteroidCollision(nums))
}

func TestB(t *testing.T) {
	nums := []int{8, -8}
	t.Log(asteroidCollision(nums))
}

func TestC(t *testing.T) {
	nums := []int{-2, -1, 1, 2}
	t.Log(asteroidCollision(nums))
}

func TestD(t *testing.T) {
	nums := []int{10, 2, -5}
	t.Log(asteroidCollision(nums))
}

func TestE(t *testing.T) {
	nums := []int{-2, -2, 1, -2}
	t.Log(asteroidCollision(nums))
}
