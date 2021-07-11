package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/number-of-provinces/
 */

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 || len(isConnected[0]) != n {
		return 0
	}
	visited := make([]bool, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs2(isConnected, visited, i)
			cnt++
		}
	}
	return cnt
}

func bfs2(isConnected [][]int, visited []bool, i int) {
	visited[i] = true
	queue := []int{i}
	for len(queue) > 0 {
		k := queue[0]
		for j := 0; j < len(isConnected); j++ {
			if isConnected[k][j] == 1 && !visited[j] {
				queue = append(queue, j)
				visited[j] = true
			}
		}
		queue = queue[1:]
	}
	return
}
