package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/
 */

type Node struct {
	i, j int
}

func ShortestPathBinaryMatrix(grid [][]int) int {
	m := len(grid)
	if m == 0 || len(grid[0]) != m || grid[0][0] != 0 || grid[m-1][m-1] != 0 {
		return -1
	}
	if m == 1 {
		return 1
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, m)
	}
	level, curLevelNode, nextLevelNode := 1, 1, 0
	queue := []Node{{0, 0}}
	visited[0][0] = true
	for len(queue) > 0 {
		i, j := queue[0].i, queue[0].j
		offset := []int{-1, 0, 1}
		for _, offI := range offset {
			for _, offJ := range offset {
				newI, newJ := i+offI, j+offJ
				if newI >= 0 && newJ >= 0 && newI < m && newJ < m && !visited[newI][newJ] && grid[newI][newJ] == 0 {
					if newI == m-1 && newJ == m-1 {
						return level + 1
					}
					queue = append(queue, Node{newI, newJ})
					nextLevelNode++
					visited[newI][newJ] = true
				}
			}
		}
		curLevelNode--
		if curLevelNode == 0 {
			level++
			curLevelNode, nextLevelNode = nextLevelNode, 0
		}
		queue = queue[1:]
	}
	return -1
}
