package _92__Number_of_Matching_Subsequences

import "sort"

func numMatchingSubseq(S string, words []string) int {
	pos := make(map[rune][]int, len(S))
	for i, b := range S {
		if _, ok := pos[b]; !ok {
			pos[b] = make([]int, 0, 32)
		}
		pos[b] = append(pos[b], i)
	}

	var matchedNum int
	for _, word := range words {
		isValid := true
		prev := -1
		for _, b := range word {
			if _, ok := pos[b]; !ok {
				isValid = false
				break
			}
			prev = searchPos(pos[b], prev)
			if prev < 0 {
				isValid = false
				break
			}
		}
		if isValid {
			matchedNum++
		}
	}

	return matchedNum
}

func searchPos(pos []int, prev int) int {
	f := func(i int) bool {
		if pos[i] <= prev {
			return false
		}
		return true
	}
	n := len(pos)
	i := sort.Search(n, f)
	if i == n {
		return -1
	}
	return pos[i]
}
