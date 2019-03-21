package _92__Number_of_Matching_Subsequences

import "testing"

func TestA(t *testing.T) {
	if 3 != numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}) {
		t.Fail()
	}
}

func TestB(t *testing.T) {
	if 2 != numMatchingSubseq("edfabc", []string{"dbc", "afb", "dac", "ace"}) {
		t.Fail()
	}
}

func TestC(t *testing.T) {
	if 0 != numMatchingSubseq("edfabc", []string{}) {
		t.Fail()
	}
}
