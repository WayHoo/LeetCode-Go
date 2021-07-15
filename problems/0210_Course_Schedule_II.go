package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/course-schedule-ii/
 */

// FindOrder 拓扑排序 + BFS
func FindOrder(numCourses int, prerequisites [][]int) []int {
	order := make([]int, 0)
	edges := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	queue := make([]int, 0)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		edges[a] = append(edges[a], b)
		indegrees[b]++
	}
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		a := queue[0]
		order = append(order, a)
		for _, b := range edges[a] {
			if indegrees[b]--; indegrees[b] == 0 {
				queue = append(queue, b)
			}
		}
		queue = queue[1:]
	}
	if len(order) != numCourses {
		return nil
	}
	for i, j := 0, numCourses-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}

// FindOrderV2 DFS
func FindOrderV2(numCourses int, prerequisites [][]int) []int {
	order := make([]int, 0)
	edges := make([][]int, numCourses)
	visited := make([]int8, numCourses)
	var dfs func(n int) bool
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		edges[a] = append(edges[a], b)
	}
	dfs = func(n int) bool {
		if visited[n] == 1 {
			// 此轮 dfs 已访问，构成环路
			return false
		}
		if visited[n] == -1 {
			// 从其他节点开始的 dfs 已访问
			return true
		}
		visited[n] = 1
		for _, m := range edges[n] {
			if !dfs(m) {
				return false
			}
		}
		visited[n] = -1
		order = append(order, n)
		return true
	}
	for i := 0; i < numCourses && dfs(i); i++ {}
	if len(order) != numCourses {
		order = nil
	}
	return order
}
