package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/decode-ways/
 */

func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	prePre, pre := 1, 1
	for i := 1; i < len(s); i++ {
		cur := 0
		if s[i] > '0' {
			cur = pre
		}
		if s[i-1] != '0' && (s[i-1]-'0')*10+(s[i]-'0') <= 26 {
			cur += prePre
		}
		if cur == 0 {
			return 0
		}
		prePre, pre = pre, cur
	}
	return pre
}
