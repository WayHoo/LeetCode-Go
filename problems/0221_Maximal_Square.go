package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/maximal-square/
 */

func MaximalSquare(matrix [][]byte) int {
	min := func(nums ...int) int {
		ans := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < ans {
				ans = nums[i]
			}
		}
		return ans
	}
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		dp[i][0] = int(matrix[i][0] - '0')
	}
	for j := 0; j < len(matrix[0]); j++ {
		dp[0][j] = int(matrix[0][j] - '0')
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
		}
	}
	maxLen := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen * maxLen
}
