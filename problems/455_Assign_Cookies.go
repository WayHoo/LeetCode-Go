package problems

import "sort"

/*
 LeetCode: https://leetcode-cn.com/problems/assign-cookies/description/
*/

func FindContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	cnt := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			cnt++
			i++
		}
		j++
	}
	return cnt
}
