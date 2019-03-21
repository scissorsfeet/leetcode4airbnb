package a_Round_Numbers

import (
	"testing"
)

func TestA(t *testing.T) {
	nums := []float64{30.3, 2.4, 3.5}
	t.Log(RoundNumbers(nums))
}

func TestB(t *testing.T) {
	nums := []float64{30.9, 2.4, 3.9}
	t.Log(RoundNumbers(nums))
}

func TestC(t *testing.T) {
	nums := []float64{1.4}
	t.Log(RoundNumbers(nums))
}
