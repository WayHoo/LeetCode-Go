package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/ones-and-zeroes/
 */

func FindMaxForm(strs []string, m int, n int) int {
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		o, z := 0, 0
		for _, ch := range str {
			if ch == '1' {
				o++
			} else {
				z++
			}
		}
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
			}
		}
	}
	return dp[m][n]
}
