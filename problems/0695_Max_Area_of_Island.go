package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/max-area-of-island/
 */

func MaxAreaOfIslandBFS(grid [][]int) int {
	m := len(grid)
	if m == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid[0])
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if tmp := bfs(grid, i, j); tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}

func bfs(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	queue := [][2]int{{i, j}}
	grid[i][j] = 0
	cnt := 1
	dire := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		for _, d := range dire {
			newI, newJ := x + d[0], y + d[1]
			if newI >= 0 && newI < m && newJ >= 0 && newJ < n && grid[newI][newJ] == 1 {
				queue = append(queue, [2]int{newI, newJ})
				grid[newI][newJ] = 0
				cnt++
			}
		}
		queue = queue[1:]
	}
	return cnt
}

func MaxAreaOfIslandDFS(grid [][]int) int {
	m := len(grid)
	if m == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, max := len(grid[0]), 0
	for i := 0; i < m;i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if tmp := dfs(grid, i, j); tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}

func dfs(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	area := 1
	dire := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for idx := range dire {
		x, y := i + dire[idx][0], j + dire[idx][1]
		if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
			area += dfs(grid, x, y)
		}
	}
	return area
}
