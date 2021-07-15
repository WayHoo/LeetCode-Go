package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/redundant-connection/
 */

func FindRedundantConnection(edges [][]int) []int {
	N := len(edges)
	ufSet := make([]int, N+1)
	rank := make([]int, N+1)
	for i := 1; i <= N; i++ {
		ufSet[i] = i
		rank[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if x != ufSet[x] {
			parent := find(ufSet[x])
			if parent != ufSet[x] {
				rank[ufSet[x]]--
				ufSet[x] = parent
			}
		}
		return ufSet[x]
	}
	union := func(x, y int) bool {
		x, y = find(x), find(y)
		if x == y {
			return false
		}
		if rank[x] <= rank[y] {
			ufSet[x] = y
			rank[y] += rank[x]
		} else {
			ufSet[y] = x
			rank[x] += rank[y]
		}
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
