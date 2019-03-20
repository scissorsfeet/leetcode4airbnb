package _48__Shortest_Completing_Word

func shortestCompletingWord(licensePlate string, words []string) string {
	res := ""
	resScore := 0
	licensePlateMap := make(map[rune]int)
	buildMap(licensePlateMap, licensePlate)

	for _, word := range words {
		score := 0
		m := make(map[rune]int)
		buildMap(m, word)
		for b, _ := range m {
			if _, ok := licensePlateMap[b]; ok {
				score += min(m[b], licensePlateMap[b])
			}
		}
		if score > resScore || (score == resScore && len(word) < len(res)) {
			resScore = score
			res = word
		}
	}
	return res
}

func buildMap(m map[rune]int, word string) {
	for _, b := range word {
		if b >= 65 && b <= 90 {
			b += 32
		}
		m[b] += 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
