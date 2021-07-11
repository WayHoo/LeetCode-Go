package problems

/*
 LeetCode: https://leetcode-cn.com/problems/valid-palindrome-ii/
*/

func ValidPalindrome(s string) bool {
	if s == "" {
		return true
	}
	r := []rune(s)
	left, right := 0, len(r)-1
	for left < right {
		if r[left] == r[right] {
			left++
			right--
		} else {
			if palindrome(r[left+1:right+1]) || palindrome(r[left:right]) {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func palindrome(r []rune) bool {
	if len(r) == 0 {
		return true
	}
	left, right := 0, len(r)-1
	for left < right {
		if r[left] == r[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
