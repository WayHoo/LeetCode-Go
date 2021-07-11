package problems

/*
 LeetCode: https://leetcode-cn.com/problems/number-of-islands/
*/

func NumIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid[0])
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				bfs1(grid, i, j)
			}
		}
	}
	return cnt
}

func bfs1(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	queue := [][2]int{{i, j}}
	dire := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		for _, d := range dire {
			x, y := queue[0][0] + d[0], queue[0][1] + d[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
				queue = append(queue, [2]int{x, y})
				grid[x][y] = '0'
			}
		}
		queue = queue[1:]
	}
	return
}
