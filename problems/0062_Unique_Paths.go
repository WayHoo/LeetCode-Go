package problems

import "math/big"

/**
 * LeetCode: https://leetcode-cn.com/problems/unique-paths/
 */

// UniquePaths 动态规划方式求解
func UniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

// UniquePathsMath 数学方式求解
func UniquePathsMath(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
