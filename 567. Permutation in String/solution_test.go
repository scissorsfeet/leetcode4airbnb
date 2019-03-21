package _67__Permutation_in_String

import "testing"

func TestA(t *testing.T) {
	if !checkInclusion("ab", "eidbaooo") {
		t.Fatal()
	}
}

func TestB(t *testing.T) {
	if !checkInclusion("adc", "dcda") {
		t.Fatal()
	}
}
