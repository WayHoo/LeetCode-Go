package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/course-schedule/
 * 解题思路：https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
 */

// CanFinish BFS + 拓扑排序
func CanFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		edges[a] = append(edges[a], b)
		indegrees[b]++
	}
	queue := make([]int, 0)
	for n, val := range indegrees {
		if val == 0 {
			queue = append(queue, n)
		}
	}
	courseCnt := 0
	for len(queue) > 0 {
		courseCnt++
		for _, val := range edges[queue[0]] {
			if indegrees[val]--; indegrees[val] == 0 {
				queue = append(queue, val)
			}
		}
		queue = queue[1:]
	}
	return courseCnt == numCourses
}

// CanFinishV2 DFS解法
func CanFinishV2(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		edges[a] = append(edges[a], b)
	}
	visited := make([]int8, numCourses)
	var dfs func(n int) bool
	dfs = func(n int) bool {
		if visited[n] == -1 {
			// 被其他节点开始的 dfs 访问过
			return true
		}
		if visited[n] == 1 {
			// 被本轮的 dfs 访问过，形成环路
			return false
		}
		visited[n] = 1
		for _, m := range edges[n] {
			if !dfs(m) {
				return false
			}
		}
		visited[n] = -1
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
