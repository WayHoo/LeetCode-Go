package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
 */

func KthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := l + (r - l) / 2
		count := 0
		for i, j := n-1, 0; i >= 0 && j < n; {
			if matrix[i][j] <= m {
				count += i + 1
				j++
			} else {
				i--
			}
		}
		if count >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
