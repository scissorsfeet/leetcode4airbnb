package _01__Online_Stock_Span

type StockSpanner struct {
	prices []int
	span   []int
	cursor int
}

func Constructor() StockSpanner {
	return StockSpanner{make([]int, 0, 1500), make([]int, 0, 1500), 0}
}

func (this *StockSpanner) Next(price int) int {
	this.prices = append(this.prices, price)
	p := this.cursor - 1
	span := 1
	for p >= 0 && price >= this.prices[p] {
		span += this.span[p]
		p -= this.span[p]
	}
	this.span = append(this.span, span)
	this.cursor++
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
