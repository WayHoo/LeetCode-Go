package problems

/*
 LeetCode: https://leetcode-cn.com/problems/is-subsequence/
*/

// IsSubsequence 双指针解法
func IsSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if t[j] == s[i] {
			i++
		}
		j++
	}
	return i == m
}

// IsSubsequenceDP 我悟了，万物皆DP
func IsSubsequenceDP(s string, t string) bool {
	m, n := len(s), len(t)
	dp := make([][26]int, n+1)
	for i := 0; i < 26; i++ {
		dp[n][i] = n
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == uint8(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	for i, idx := 0, 0; i < m; i++ {
		j := s[i] - 'a'
		if dp[idx][j] == n {
			return false
		}
		idx = dp[idx][j] + 1
	}
	return true
}
