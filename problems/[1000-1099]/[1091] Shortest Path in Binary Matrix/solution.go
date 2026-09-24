package main

func shortestPathBinaryMatrixBFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[m-1][n-1] == 1 {
		return -1
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{{0, 0}}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	length := 1

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr[0] == m-1 && curr[1] == n-1 {
				return length
			}

			for _, dir := range dirs {
				next := []int{curr[0] + dir[0], curr[1] + dir[1]}
				if next[0] >= 0 && next[0] < m && next[1] >= 0 && next[1] < n &&
					!visited[next[0]][next[1]] &&
					grid[next[0]][next[1]] == 0 {
					visited[next[0]][next[1]] = true
					queue = append(queue, next)
				}
			}
		}

		length++
	}

	return length
}

func shortestPathBinaryMatrixDFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[m-1][n-1] == 1 {
		return -1
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	minLength := -1

	var dfs func(curr []int, length int)
	dfs = func(curr []int, length int) {
		if curr[0] == m-1 && curr[1] == n-1 {
			if minLength == -1 || length < minLength {
				minLength = length
			}

			return
		}

		if minLength != -1 && length >= minLength {
			return
		}

		visited[curr[0]][curr[1]] = true

		for _, dir := range dirs {
			next := []int{curr[0] + dir[0], curr[1] + dir[1]}
			if next[0] >= 0 && next[0] < m && next[1] >= 0 && next[1] < n &&
				!visited[next[0]][next[1]] &&
				grid[next[0]][next[1]] == 0 {
				dfs(next, length+1)
			}
		}

		visited[curr[0]][curr[1]] = false
	}

	dfs([]int{0, 0}, 1)

	return minLength
}

func main() {
	grid := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}

	res1 := shortestPathBinaryMatrixBFS(grid)
	res2 := shortestPathBinaryMatrixDFS(grid)
	print(res1, res2)
}
