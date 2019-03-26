package _69_Alien_Dictionary

import "testing"

func TestA(t *testing.T) {
	words := []string{
		"wrt",
		"wrf",
		"er",
		"ett",
		"rftt",
	}

	t.Log(alienOrders(words))
}
