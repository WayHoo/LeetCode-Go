package problems

/*
 LeetCode: https://leetcode-cn.com/problems/surrounded-regions/
*/

func Solve(board [][]byte) {
	m := len(board)
	if m == 0 || len(board[0]) == 0 {
		return
	}
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				bfs3(board, visited, i, j)
			}
		}
	}
	return
}

func bfs3(board [][]byte, visited [][]bool, i, j int) {
	m, n := len(board), len(board[0])
	queue := [][2]int{{i, j}}
	list, fill := [][2]int{{i, j}}, true
	if isOnBorder(m, n, i, j) {
		fill = false
	}
	visited[i][j] = true
	dire := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		for _, d := range dire {
			x, y := queue[0][0] + d[0], queue[0][1] + d[1]
			if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' && !visited[x][y] {
				visited[x][y] = true
				queue = append(queue, [2]int{x, y})
				if isOnBorder(m, n, x, y) {
					fill = false
				}
				if fill {
					list = append(list, [2]int{x, y})
				}
			}
		}
		queue = queue[1:]
	}
	if fill {
		for _, cor := range list {
			x, y := cor[0], cor[1]
			board[x][y] = 'X'
		}
	}
	return
}

func isOnBorder(m, n, i, j int) bool {
	if i == 0 || i == m-1 || j == 0 || j == n-1 {
		return true
	}
	return false
}
