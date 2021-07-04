package problems

/*
 LeetCode: https://leetcode-cn.com/problems/is-subsequence/description/
*/

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
