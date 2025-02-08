package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if !dfs(i, graph, visited) {
			return false
		}
	}
	return true
}

func dfs(i int, graph [][]int, visited []int) bool {
	if visited[i] == 1 {
		return false
	}
	if visited[i] == 2 {
		return true
	}
	visited[i] = 1
	for _, j := range graph[i] {
		if !dfs(j, graph, visited) {
			return false
		}
	}
	visited[i] = 2
	return true
}