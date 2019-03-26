package _69_Alien_Dictionary

import "math"

const (
	VISITED  = 1
	VISITING = 2
)

func alienOrders(words []string) string {
	wordsLen := len(words)
	graph := make([][]int, 26)
	for i, _ := range graph {
		graph[i] = make([]int, 26)
	}
	for _, word := range words {
		for _, b := range word {
			graph[b-'a'][b-'a'] = 1
		}
	}

	for i := 0; i < wordsLen-1; i++ {
		mn := int(math.Min(float64(len(words[i])), float64(len(words[i+1]))))
		j := 0
		for ; j < mn; j++ {
			if words[i][j] != words[i+1][j] {
				graph[words[i][j]-'a'][words[i+1][j]-'a'] = 1
				break
			}
		}
		if j == mn && len(words[i]) > len(words[i+1]) {
			return ""
		}
	}

	res := ""
	note := make([]int, 26)
	for i := 0; i < 26; i++ {
		if graph[i][i] == 0 {
			continue
		}
		if dfs(i, graph, note, &res) {
			return ""
		}
	}

	resReverse := []rune(res)
	for i, j := 0, len(resReverse)-1; i < j; i, j = i+1, j-1 {
		resReverse[i], resReverse[j] = resReverse[j], resReverse[i]
	}

	return string(resReverse)
}

func dfs(i int, graph [][]int, note []int, res *string) bool {
	if note[i] == VISITED {
		return false
	}
	if note[i] == VISITING {
		return true
	}

	note[i] = VISITING

	for j := 0; j < 26; j++ {
		if i != j && graph[i][j] == 1 {
			if dfs(j, graph, note, res) {
				return true
			}
		}
	}

	note[i] = VISITED
	*res += string(i + 'a')

	return false
}
