package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/integer-break/
 */

func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		cur := 0
		for j := 1; j < i; j++ {
			tmp1, tmp2 := j*(i-j), j*dp[i-j]
			cur = max(max(cur, tmp1), tmp2)
		}
		dp[i] = cur
	}
	return dp[n]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
