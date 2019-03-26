package _87__Cheapest_Flights_Within_K_Stops

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	var cheapestPrice int = math.MaxInt32
	var visited = make([]int, n)
	var graph = make([][][]int, n)
	for i := range flights {
		src := flights[i][0]
		dst := flights[i][1]
		price := flights[i][2]
		if graph[src] == nil {
			graph[src] = make([][]int, 0)
		}
		tmp := []int{dst, price}
		graph[src] = append(graph[src], tmp)
	}
	dfs(src, dst, K, graph, visited, &cheapestPrice, 0)

	if cheapestPrice == math.MaxInt32 {
		return -1
	}
	return cheapestPrice
}

func dfs(cur, dst, K int, graph [][][]int, visited []int, pCheapestPrice *int, out int) {
	if cur == dst {
		*pCheapestPrice = out
		return
	}
	if K < 0 {
		return
	}

	for _, m := range graph[cur] {
		if visited[m[0]] == 1 || out+m[1] > *pCheapestPrice {
			continue
		}
		visited[m[0]] = 1
		dfs(m[0], dst, K-1, graph, visited, pCheapestPrice, out+m[1])
		visited[m[0]] = 0
	}
}
