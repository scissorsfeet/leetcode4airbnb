package _01__Online_Stock_Span

import "testing"

func TestA(t *testing.T) {
	s := StockSpanner{}
	t.Log(s.Next(100))
	t.Log(s.Next(80))
	t.Log(s.Next(60))
	t.Log(s.Next(70))
	t.Log(s.Next(60))
	t.Log(s.Next(75))
	t.Log(s.Next(85))
}

func TestB(t *testing.T) {
	s := StockSpanner{}
	t.Log(s.Next(50))
	t.Log(s.Next(40))
	t.Log(s.Next(50))
	t.Log(s.Next(50))
	t.Log(s.Next(30))
	t.Log(s.Next(60))
	t.Log(s.Next(50))
}
