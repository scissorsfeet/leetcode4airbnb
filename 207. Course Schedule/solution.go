package _07__Course_Schedule

const (
	VISTED   = 1
	VISITING = 2
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	note := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		hasCycle := dfs(i, graph, note)
		if hasCycle {
			return false
		}
	}

	return true
}

func dfs(i int, graph [][]int, note []int) bool {
	if note[i] == VISITING {
		return true
	}
	if note[i] == VISTED {
		return false
	}

	note[i] = VISITING

	for _, neighbour := range graph[i] {
		if dfs(neighbour, graph, note) {
			return true
		}
	}

	note[i] = VISTED
	return false
}
