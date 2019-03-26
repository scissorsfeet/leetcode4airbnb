package _55_Pour_Water

import "fmt"

func pourWater(heights []int, K, V int) {
	n := len(heights)
	srcHeights := make([]int, n)
	copy(srcHeights, heights)
	warters := make([][]int, n)
	for i := 0; i < V; i++ {
		l, r := K, K
		for l > 0 && heights[l] >= heights[l-1] {
			l--
		}
		for l < K && heights[l] == heights[l+1] {
			l++
		}
		for r < n-1 && heights[r] >= heights[r+1] {
			r++
		}
		for r > K && heights[r-1] == heights[r] {
			r--
		}
		if warters[i] == nil {
			warters[i] = make([]int, 0, 8)
		}
		if heights[l] < heights[K] {
			warters[l] = append(warters[l], i+1)
			heights[l]++
		} else {
			warters[r] = append(warters[r], i+1)
			heights[r]++
		}
	}
	fmt.Println(heights)

	maxHeight := 0
	for i := 0; i < n; i++ {
		if heights[i] > maxHeight {
			maxHeight = heights[i]
		}
	}
	fmt.Println(maxHeight, srcHeights)

	for i := maxHeight; i > 0; i-- {
		for j := 0; j < n; j++ {
			if i > heights[j] {
				fmt.Print(" ")
			} else if i <= srcHeights[j] {
				fmt.Print("*")
			} else {
				fmt.Print(warters[j][i-srcHeights[j]-1])
			}
		}
		fmt.Println()
	}
}
