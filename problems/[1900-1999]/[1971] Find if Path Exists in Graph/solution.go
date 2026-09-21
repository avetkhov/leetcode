package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)

	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	queue := []int{source}
	visited := make([]bool, n)
	visited[source] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == destination {
			return true
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	visited[source] = true

	return dfs(source, destination, graph, visited)
}

func dfs(current int, destination int, graph [][]int, visited []bool) bool {
	if current == destination {
		return true
	}

	visited[current] = true

	for _, neighbor := range graph[current] {
		if !visited[neighbor] && dfs(neighbor, destination, graph, visited) {
			return true
		}
	}

	return false
}
