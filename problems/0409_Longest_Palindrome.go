package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/longest-palindrome/
 */

func LongestPalindrome(s string) int {
	charMap := make(map[uint8]uint16)
	for _, ch := range s {
		charMap[uint8(ch)]++
	}
	count := uint16(0)
	for _, cnt := range charMap {
		if cnt%2 == 0 {
			count += cnt
		} else {
			count += cnt - 1
		}
	}
	if int(count) < len(s) {
		count++
	}
	return int(count)
}
