package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/is-graph-bipartite/
 */

func IsBipartite(graph [][]int) bool {
	tagMap := make(map[int]int8)
	visited := make([]bool, len(graph))
	var queue []int
	for i := 0; i < len(graph); i++ {
		if visited[i] {
			continue
		}
		queue, visited[i], tagMap[i] = []int{i}, true, 1
		for len(queue) > 0 {
			k := queue[0]
			for j := 0; j < len(graph[k]); j++ {
				idx := graph[k][j]
				if tagMap[idx] == tagMap[k] {
					return false
				}
				if !visited[idx] {
					visited[idx] = true
					tagMap[idx] = -tagMap[k]
					queue = append(queue, idx)
				}
			}
			queue = queue[1:]
		}
	}
	return true
}
