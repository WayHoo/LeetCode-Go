package problems

/*
 LeetCode: https://leetcode-cn.com/problems/word-search/
*/

func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dire := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word) - 1 {
			return true
		}
		visited[i][j] = true
		defer func() {visited[i][j] = false}()
		for _, d := range dire {
			x, y := i + d[0], j + d[1]
			if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
				if dfs(x, y, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
