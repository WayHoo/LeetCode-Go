package problems

import "sort"

/*
 LeetCode: https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
*/

func FindMinArrowShots(points [][]int) int {
	N := len(points)
	if N == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	cnt, end := 1, points[0][1]
	for i := 1; i < N; i++ {
		if points[i][0] > end {
			cnt++
			end = points[i][1]
		}
	}
	return cnt
}
