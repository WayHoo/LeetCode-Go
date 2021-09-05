package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/longest-common-subsequence/
 */

func LongestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, l2+1)
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}
