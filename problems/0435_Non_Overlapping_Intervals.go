package problems

import (
	"sort"
)

/*
 LeetCode: https://leetcode-cn.com/problems/assign-cookies/
*/

func EraseOverlapIntervals(intervals [][]int) int {
	N := len(intervals)
	if N == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	cnt := 1
	preEnd := intervals[0][1]
	for i := 1; i < N; i++ {
		if intervals[i][0] >= preEnd {
			cnt++
			preEnd = intervals[i][1]
		}
	}
	return N - cnt
}
