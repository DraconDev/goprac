package leetcode

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	adjList := make(map[int][]int)
	for _, e := range edges {
		adjList[e[0]] = append(adjList[e[0]], e[1])
		adjList[e[1]] = append(adjList[e[1]], e[0])
	}
	visited := make(map[int]bool)
	var dfs func(v int) bool
	dfs = func(v int) bool {
		if v == destination {
			return true
		}
		visited[v] = true
		for _, neighbor := range adjList[v] {
			if !visited[neighbor] && dfs(neighbor) {
				return true
			}
		}
		return false
	}
	return dfs(source)
}
