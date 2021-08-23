package problems

import "math"

/**
 * LeetCode: https://leetcode-cn.com/problems/edit-distance/
 */

func minDistance(word1 string, word2 string) int {
	var min func(list ...int) int
	min = func(list ...int) int {
		minVal := math.MaxInt32
		for _, val := range list {
			if val < minVal {
				minVal = val
			}
		}
		return minVal
	}
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[l1][l2]
}
