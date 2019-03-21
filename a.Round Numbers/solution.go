package a_Round_Numbers

import (
	"math"
	"sort"
)

type PairItem struct {
	index int
	value float64
}

func RoundNumbers(nums []float64) []float64 {
	var floorSum float64 = 0
	var allSum float64 = 0
	pair := make([]PairItem, 0, len(nums))
	res := make([]float64, 0, len(nums))
	for i, num := range nums {
		tmpFloor := math.Floor(num)
		res = append(res, tmpFloor)
		pair = append(pair, PairItem{i, num})
		floorSum += tmpFloor
		allSum += num
	}

	remain := int(math.Round(allSum) - floorSum)
	sort.Sort(pairItemSlice(pair))

	pos := len(pair) - 1
	for i := 1; i <= remain; i++ {
		res[pair[pos].index] += 1
		pos--
	}

	return res
}

type pairItemSlice []PairItem

func (p pairItemSlice) Len() int { return len(p) }
func (p pairItemSlice) Less(i, j int) bool {
	return p[i].value-math.Floor(p[i].value) < p[j].value-math.Floor(p[j].value)
}
func (p pairItemSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
