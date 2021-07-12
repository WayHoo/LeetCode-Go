package problems

import "fmt"

/**
 * LeetCode: https://leetcode-cn.com/problems/palindrome-number/
 */

// IsPalindrome 将整数转成字符串的解法
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strs := fmt.Sprintf("%d", x)
	for i, j := 0, len(strs)-1; i <= j; i, j = i+1, j-1 {
		if strs[i] != strs[j] {
			return false
		}
	}
	return true
}

// IsPalindromeV2 将整数分成左右两部分，右边那部分需要转置，然后判断这两部分是否相等。
func IsPalindromeV2(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	right := 0
	for right < x {
		right = right * 10 + x % 10
		x /= 10
	}
	return right == x || x == right/10
}
