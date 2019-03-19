package _28__Longest_Consecutive_Sequence

type Item struct {
	l int
	r int
}

func longestConsecutive(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	var res int
	hm := make(map[int]*Item, numsLen)
	for _, num := range nums {
		if _, ok := hm[num]; ok {
			continue
		}
		_, okl := hm[num-1]
		_, okr := hm[num+1]
		tmp := 0
		var l, r int
		if okl && okr {
			l = hm[num-1].l
			r = hm[num+1].r
		} else if okl {
			l = hm[num-1].l
			r = num
		} else if okr {
			l = num
			r = hm[num+1].r
		} else {
			l = num
			r = num
		}
		hm[num] = &Item{l, r}
		hm[l].r = r
		hm[r].l = l
		tmp = r - l + 1
		if tmp > res {
			res = tmp
		}
	}
	return res
}
