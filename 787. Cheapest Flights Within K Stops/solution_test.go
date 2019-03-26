package _87__Cheapest_Flights_Within_K_Stops

import "testing"

func TestA(t *testing.T) {
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	p := findCheapestPrice(3, flights, 0, 2, 0)
	if p != 500 {
		t.Fatal(p)
	}
}

func TestB(t *testing.T) {
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	p := findCheapestPrice(3, flights, 0, 2, 2)
	if p != 200 {
		t.Fatal(p)
	}
}
