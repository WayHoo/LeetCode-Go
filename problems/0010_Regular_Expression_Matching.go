package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/regular-expression-matching/
 */

// IsMatchDP DP方法实现的正则匹配
func IsMatchDP(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j == 0 {
				dp[i][j] = i == 0
				continue
			}
			if p[j-1] == '*' {
				// 匹配0次
				if j > 1 {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
				// 匹配多次
				if i > 0 && j > 1 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
