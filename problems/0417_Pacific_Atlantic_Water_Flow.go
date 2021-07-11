package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
 解题思路：以左边界和上边界的元素为起点做“水往高处流”的 bfs，遍历到的元素标记为可流入太平洋；
		 以右边界和下边界的元素为起点做同样的 bfs，遍历到的元素标记为可流入大西洋；
		 遍历标记数组，有上述两种标记的坐标点即为解空间元素。
 */

func PacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	if m == 0 || len(heights[0]) == 0 {
		return nil
	}
	n := len(heights[0])
	reach := make([][]int8, m)
	visitP, visitA := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		reach[i] = make([]int8, n)
		visitP[i] = make([]bool, n)
		visitA[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		bfs4(heights, reach, visitP, 0, j)
		bfs4(heights, reach, visitA, m-1, j)
	}
	for i := 0; i < m; i++ {
		bfs4(heights, reach, visitP, i, 0)
		bfs4(heights, reach, visitA, i, n-1)
	}
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if reach[i][j] == 2 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs4(heights [][]int, reach [][]int8, visited [][]bool, i, j int) {
	if visited[i][j] {
		return
	}
	m, n := len(heights), len(heights[0])
	reach[i][j]++
	queue := [][2]int{{i, j}}
	visited[i][j] = true
	dire := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		a, b := queue[0][0], queue[0][1]
		cur := heights[a][b]
		for _, d := range dire {
			x, y := a+d[0], b+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && heights[x][y] >= cur && !visited[x][y] {
				visited[x][y] = true
				queue = append(queue, [2]int{x, y})
				reach[x][y]++
			}
		}
		queue = queue[1:]
	}
}
