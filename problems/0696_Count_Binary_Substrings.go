package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/count-binary-substrings/
 */

func CountBinarySubstrings(s string) int {
	preLen, curLen, cnt := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curLen++
		} else {
			preLen, curLen = curLen, 1
		}
		if curLen <= preLen {
			cnt++
		}
	}
	return cnt
}
