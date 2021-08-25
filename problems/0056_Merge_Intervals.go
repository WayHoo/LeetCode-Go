package problems

import "sort"

/**
 * LeetCode: https://leetcode-cn.com/problems/merge-intervals/
 */

func MergeIntervals(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] <= intervals[j][0]
    })
    max := func(a, b int) int {
        if a >= b {
            return a
        }
        return b
    }
    ans := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        pre := ans[len(ans)-1]
        cur := intervals[i]
        if cur[0] <= pre[1] {
            merged := []int{pre[0], max(pre[1], cur[1])}
            ans[len(ans)-1] = merged
        } else {
            ans = append(ans, cur)
        }
    }
    return ans
}
