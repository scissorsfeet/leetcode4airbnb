package _10__Course_Schedule_II

const (
	VISTED   = 1
	VISITING = 2
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}
	note := make([]int, numCourses)
	order := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if dfs(i, graph, note, &order) {
			return []int{}
		}
	}
	for i, j := 0, numCourses-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}

func dfs(i int, graph [][]int, note []int, order *[]int) bool {
	if note[i] == VISTED {
		return false
	}
	if note[i] == VISITING {
		return true
	}

	note[i] = VISITING

	for _, neighbour := range graph[i] {
		if dfs(neighbour, graph, note, order) {
			return true
		}
	}

	note[i] = VISTED
	*order = append(*order, i)
	return false
}
