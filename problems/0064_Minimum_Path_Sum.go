package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/minimum-path-sum/
 */

func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				// 只能从上面过来
				dp[j] = dp[j]
			} else if i == 0 {
				// 只能从左侧过来
				dp[j] = dp[j-1]
			} else if dp[j-1] < dp[j] {
				dp[j] = dp[j-1]
			}
			dp[j] += grid[i][j]
		}
	}
	return dp[n-1]
}
